package repository

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
)

type TokenRepository interface {
	RegisterToken(ctx context.Context, token domain.Token) error
	UnregisterToken(ctx context.Context, token string) error
	LoadTokens(ctx context.Context) ([]*domain.Token, error)
	ValidateToken(ctx context.Context, token string) (bool, error)
}
