package main

import (
	"flag"
	"github.com/Siroshun09/plugin-list/app"
	"github.com/Siroshun09/plugin-list/repository/sqlite"
	"log"
	"log/slog"
	"os"
)

func main() {
	slog.Info("Starting plugin-list...")

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
	a, err := app.NewApp(conn)

	if err != nil {
		slog.Error("Failed to initialize the application", err)
		os.Exit(1)
	}

	port := *flag.String("port", "8080", "The port to listen on")
	slog.Info("Preparing the server...", slog.String("port", port))
	err = a.PrepareServer(port)

	if err != nil {
		slog.Error("Failed to initialize the web server", err)
		os.Exit(1)
		return
	}

	slog.Info("The server has started!")
	if err = a.Start(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
