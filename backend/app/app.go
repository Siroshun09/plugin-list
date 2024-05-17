package app

import (
	"context"
	"errors"
	"github.com/Siroshun09/plugin-list/api"
	"github.com/Siroshun09/plugin-list/repository/sqlite"
	"github.com/Siroshun09/plugin-list/usecase"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	middleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

// App は API サーバーを稼働させるのに必要なものを保持し、API サーバーを起動・終了するメソッドを提供します。
type App struct {
	// domain.MCPlugin を取得・保存・削除するための usecase.MCPluginUseCase
	McPluginUseCase usecase.MCPluginUseCase
	// domain.Token を取得するための usecase.TokenUseCase
	TokenUseCase usecase.TokenUseCase
	// http.Server インスタンス
	Server *http.Server
}

var (
	errInvalidSecuritySchema = errors.New("invalid Security Schema")
)

// NewApp は sqlite.Connection を使用して新しい App を作成します。
// このメソッドでは、データベースへのテーブル作成も行われます。
// App.Server() はこの段階では初期化されておらず、nil を返します。
func NewApp(ctx context.Context, conn sqlite.Connection) (*App, error) {
	slog.Info("Initializing the repository for MCPlugins...")
	mcPluginRepo, err := conn.NewMCPluginRepository(ctx)

	if err != nil {
		slog.Error("Failed to initialize the repository for MCPlugins", err)
		os.Exit(1)
	}

	tokenRepo, err := conn.NewTokenRepository(ctx)

	if err != nil {
		slog.Error("Failed to initialize the repository for tokens", err)
		os.Exit(1)
	}

	mcPluginUseCase := usecase.NewMCPluginUseCase(mcPluginRepo)
	tokenUseCase := usecase.NewTokenUseCase(tokenRepo)

	return &App{mcPluginUseCase, tokenUseCase, nil}, nil
}

// PrepareServer は指定されたポート番号を使用して API サーバーを起動する準備を行います。
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

	r.Use(cors.AllowAll().Handler) // FIXME

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

// Start は API サーバーを起動し、通信を待機します。
func (app *App) Start() {
	if err := app.Server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

// HandleShutdown は渡された context.Context が完了した際に、API サーバーのシャットダウンを行います。
func (app *App) HandleShutdown(ctx context.Context) error {
	<-ctx.Done()
	slog.Info("Stopping the web server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return app.Server.Shutdown(ctx)
}
