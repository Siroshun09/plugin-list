package sqlite

import (
	"context"
	"database/sql"
	"errors"
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

	insertTokenQuery = "INSERT INTO tokens (token, created_at) VALUES (?, ?)"

	deleteTokenQuery = "DELETE FROM tokens WHERE token = ?"

	selectAllTokenQuery = "SELECT token, created_at FROM tokens"

	validateTokenQuery = "SELECT COUNT(token) FROM tokens WHERE token = ?"
)

func (c *sqliteConnection) NewTokenRepository(ctx context.Context) (repository.TokenRepository, error) {
	if _, err := c.db.ExecContext(ctx, tokenTableSchema); err != nil {
		return nil, err
	}

	return tokenRepository{c}, nil
}

func (t tokenRepository) AddToken(ctx context.Context, token domain.Token) (returnErr error) {
	rows, err := t.conn.db.QueryContext(ctx, insertTokenQuery, token.Value, token.Created.UnixMilli())

	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	return nil
}

func (t tokenRepository) RemoveToken(ctx context.Context, token string) (returnErr error) {
	rows, err := t.conn.db.QueryContext(ctx, deleteTokenQuery, token)

	if err != nil {
		return err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
		}
	}(rows)

	return nil
}

func (t tokenRepository) LoadTokens(ctx context.Context) (tokens []*domain.Token, returnErr error) {
	rows, err := t.conn.db.QueryContext(ctx, selectAllTokenQuery)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
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

func (t tokenRepository) ValidateToken(ctx context.Context, token string) (valid bool, returnErr error) {
	rows, err := t.conn.db.QueryContext(ctx, validateTokenQuery, token)

	if err != nil {
		return false, err
	}

	defer func(rows *sql.Rows) {
		closeErr := rows.Close()
		if closeErr != nil {
			returnErr = errors.Join(err, closeErr)
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
