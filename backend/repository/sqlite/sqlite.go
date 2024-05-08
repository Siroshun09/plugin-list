package sqlite

import (
	"database/sql"
	"github.com/Siroshun09/plugin-list/repository"
	_ "modernc.org/sqlite"
)

const (
	ImplementationName = "sqlite"
	DatabaseFilename   = "sqlite.db"
)

type Connection interface {
	IsOpen() bool
	Close() error
	NewMCPluginRepository() (repository.MCPluginRepository, error)
	NewTokenRepository() (repository.TokenRepository, error)
}

type sqliteConnection struct {
	db *sql.DB
}

func CreateConnection(filepath string) (Connection, error) {
	db, err := sql.Open("sqlite", filepath)

	if err != nil {
		return nil, err
	}

	return &sqliteConnection{db}, nil
}

func (c *sqliteConnection) IsOpen() bool {
	return c.db.Ping() == nil
}

func (c *sqliteConnection) Close() error {
	return c.db.Close()
}
