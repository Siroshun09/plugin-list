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

// Connection は SQLite データベースとの接続を保持します。
type Connection interface {
	// IsOpen は SQLite データベースとの接続が現在も有効かを判定します。
	IsOpen() bool
	// Close は SQLite データベースとの接続を終了します。
	Close() error
	// NewMCPluginRepository はこの接続を使用した repository.MCPluginRepository を作成します。
	NewMCPluginRepository() (repository.MCPluginRepository, error)
	// NewTokenRepository はこの接続を使用した repository.TokenRepository を作成します。
	NewTokenRepository() (repository.TokenRepository, error)
}

type sqliteConnection struct {
	db *sql.DB
}

// CreateConnection は指定されたファイルパスの SQLite データベースとの接続を作成します。
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
