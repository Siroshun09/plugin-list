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
	// CreateNewRandomToken は指定された長さでランダムなバイト列を生成し、domain.Token を作成します。
	// バイト列は crypto/rand によって生成されます。
	// domain.Token の文字列トークンには、そのバイト列を16進数表記に変換したものを、作成時間には time.Now() が使用されます。
	CreateNewRandomToken(ctx context.Context, length int) (*domain.Token, error)
	// GetAllTokens は現在有効なすべてのトークンを取得します。
	GetAllTokens(ctx context.Context) ([]domain.Token, error)
	// ValidateToken は指定されたトークンの文字列が有効かどうか判定します。
	ValidateToken(ctx context.Context, token string) (bool, error)
	// InvalidateToken は指定されたトークンの文字列を無効化します。
	// トークンがすでに無効であったとしても、このメソッドはエラーを返しません。
	InvalidateToken(ctx context.Context, token string) error
}

// NewTokenUseCase は repository.TokenRepository から新しい TokenUseCase を作成します。
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

	token := domain.Token{Value: hex.EncodeToString(b), Created: time.Now()}
	if err := r.repo.AddToken(ctx, token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (r tokenRepositoryUseCase) GetAllTokens(ctx context.Context) ([]domain.Token, error) {
	return r.repo.LoadTokens(ctx)
}

func (r tokenRepositoryUseCase) ValidateToken(ctx context.Context, token string) (bool, error) {
	return r.repo.ValidateToken(ctx, token)
}

func (r tokenRepositoryUseCase) InvalidateToken(ctx context.Context, token string) error {
	return r.repo.RemoveToken(ctx, token)
}
