package app

import (
	"context"
	"errors"
	"github.com/Siroshun09/plugin-list/api"
	"github.com/Siroshun09/plugin-list/repository/sqlite"
	"github.com/Siroshun09/plugin-list/usecase"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	McPluginUseCase usecase.MCPluginUseCase
	TokenUseCase    usecase.TokenUseCase
	Server          *http.Server
}

var (
	errInvalidSecuritySchema = errors.New("invalid Security Schema")
)

func NewApp(conn sqlite.Connection) (*App, error) {
	slog.Info("Initializing the repository for MCPlugins...")
	mcPluginRepo, err := conn.NewMCPluginRepository()

	if err != nil {
		slog.Error("Failed to initialize the repository for MCPlugins", err)
		os.Exit(1)
	}

	tokenRepo, err := conn.NewTokenRepository()

	if err != nil {
		slog.Error("Failed to initialize the repository for tokens", err)
		os.Exit(1)
	}

	mcPluginUseCase := usecase.NewMCPluginUseCase(mcPluginRepo)
	tokenUseCase := usecase.NewTokenUseCase(tokenRepo)

	return &App{mcPluginUseCase, tokenUseCase, nil}, nil
}

func (app *App) PrepareServer(port string) error {
	// https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/chi/petstore.go
	swagger, err := api.GetSwagger()

	if err != nil {
		return err
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	pluginList := api.NewPluginList(app.McPluginUseCase)

	validatorOpts := &middleware.Options{}

	validatorOpts.Options.AuthenticationFunc = func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		switch input.SecuritySchemeName {
		case "Token":
			return api.ValidateToken(app.TokenUseCase, ctx, input)
		default:
			return errInvalidSecuritySchema
		}
	}

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	r.Use(middleware.OapiRequestValidatorWithOptions(swagger, validatorOpts))

	// We now register our plugin-list above as the handler for the interface
	api.HandlerFromMux(pluginList, r)

	app.Server = &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", port),
	}

	return nil
}

func (app *App) Start() error {
	go func() {
		if err := app.Server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.Server.Shutdown(ctx)
}
