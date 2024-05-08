package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
	"time"
)

// TokenUseCase はトークンの管理や検証を行います
type TokenUseCase interface {
	CreateNewRandomToken(ctx context.Context, length int) (*domain.Token, error)
	GetAllTokens(ctx context.Context) ([]*domain.Token, error)
	InvalidateToken(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) (bool, error)
}

func NewTokenUseCase(repo repository.TokenRepository) TokenUseCase {
	return tokenRepositoryUseCase{repo}
}

type tokenRepositoryUseCase struct {
	repo repository.TokenRepository
}

func (r tokenRepositoryUseCase) CreateNewRandomToken(ctx context.Context, length int) (*domain.Token, error) {
	b := make([]byte, length)

	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	token := domain.NewToken(hex.EncodeToString(b), time.Now())

	if err := r.repo.RegisterToken(ctx, token); err != nil {
		return nil, err
	}

	return token, nil
}

func (r tokenRepositoryUseCase) GetAllTokens(ctx context.Context) ([]*domain.Token, error) {
	return r.repo.LoadTokens(ctx)
}

func (r tokenRepositoryUseCase) InvalidateToken(ctx context.Context, token string) error {
	return r.repo.UnregisterToken(ctx, token)
}

func (r tokenRepositoryUseCase) ValidateToken(ctx context.Context, token string) (bool, error) {
	return r.repo.ValidateToken(ctx, token)
}
