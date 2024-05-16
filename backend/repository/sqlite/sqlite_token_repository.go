package sqlite

import (
	"context"
	"database/sql"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
	"time"
)

type tokenRepository struct {
	conn *sqliteConnection
}

const (
	tokenTableSchema = `
		CREATE TABLE IF NOT EXISTS tokens (
			token VARCHAR(32) NOT NULL PRIMARY KEY,
			created_at INTEGER NOT NULL
		)
		`

	registerTokenQuery = "INSERT INTO tokens (token, created_at) VALUES (?, ?)"

	unregisterTokenQuery = "DELETE FROM tokens WHERE token = ?"

	selectAllTokenQuery = "SELECT token, created_at FROM tokens"

	validateTokenQuery = "SELECT COUNT(token) FROM tokens WHERE token = ?"
)

func (c *sqliteConnection) NewTokenRepository() (repository.TokenRepository, error) {
	if _, err := c.db.Exec(tokenTableSchema); err != nil {
		return nil, err
	}

	return tokenRepository{c}, nil
}

func (t tokenRepository) RegisterToken(_ context.Context, token domain.Token) (returnErr error) {
	rows, err := t.conn.db.Query(registerTokenQuery, token.Value, token.Created.UnixMilli())

	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			returnErr = err
		}
	}(rows)

	return nil
}

func (t tokenRepository) UnregisterToken(_ context.Context, token string) (returnErr error) {
	rows, err := t.conn.db.Query(unregisterTokenQuery, token)

	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			returnErr = err
		}
	}(rows)

	return nil
}

func (t tokenRepository) LoadTokens(_ context.Context) (tokens []*domain.Token, returnErr error) {
	rows, err := t.conn.db.Query(selectAllTokenQuery)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			returnErr = err
		}
	}(rows)

	var result []*domain.Token

	for rows.Next() {
		var token domain.Token
		var createdAt int64
		if err := rows.Scan(&token.Value, &createdAt); err != nil {
			return nil, err
		}
		token.Created = time.UnixMilli(createdAt)
		result = append(result, &token)
	}

	return result, nil
}

func (t tokenRepository) ValidateToken(_ context.Context, token string) (valid bool, returnErr error) {
	rows, err := t.conn.db.Query(validateTokenQuery, token)

	if err != nil {
		return false, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			returnErr = err
		}
	}(rows)

	rows.Next()

	var count int

	err = rows.Scan(&count)
	if err != nil {
		return false, err
	}

	return count != 0, nil
}
