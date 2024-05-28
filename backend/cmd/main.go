package main

import (
	"context"
	"github.com/Siroshun09/plugin-list/app"
	"github.com/Siroshun09/plugin-list/handler"
	"github.com/Siroshun09/plugin-list/repository/sqlite"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	slog.Info("Starting plugin-list...")

	app.ParseFlags()

	dbPath := "./" + sqlite.DatabaseFilename
	slog.Info("Initializing the database...", slog.String("implementation", sqlite.ImplementationName), slog.Any("db_path", dbPath))
	conn, err := sqlite.CreateConnection(dbPath)

	if err != nil {
		slog.Error("Failed to initialize the database", err)
		os.Exit(1)
	}

	defer func(conn sqlite.Connection) {
		err := conn.Close()
		if err != nil {
			slog.Error("Failed to close the database", err)
			os.Exit(1)
		}
	}(conn)

	slog.Info("Initializing the application...")
	a, err := app.NewApp(context.Background(), conn)

	if err != nil {
		slog.Error("Failed to initialize the application", err)
		os.Exit(1)
	}

	slog.Info(
		"Preparing the server...",
		slog.String(app.PortFlag, app.GetPort()),
		slog.String(app.AllowedOriginFlag+"s", app.GetAllowedOriginsAsString()),
		slog.Bool(app.PrintUnknownOriginsFlag, app.PrintUnknownOrigins()),
	)
	err = a.PrepareServer(app.GetPort(), app.GetAllowedOrigins(), app.PrintUnknownOrigins())

	if err != nil {
		slog.Error("Failed to initialize the web server", err)
		os.Exit(1)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		quit, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt, os.Kill)
		defer stop()

		<-quit.Done()
		cancel()
	}()

	go func() {
		err := a.HandleShutdown(ctx)
		if err != nil {
			slog.Error("Failed to shutdown the web server", err)
			os.Exit(1)
		}
	}()

	go a.Start()

	slog.Info("The server has started!")

	go handler.HandleConsoleInput(a.TokenUseCase, cancel)

	<-ctx.Done()
	slog.Info("Stopping plugin-list...")
}
