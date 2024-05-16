package repository

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
)

// TokenRepository は domain.Token のデータ管理を行います。
type TokenRepository interface {
	// AddToken は新しい domain.Token を登録します。
	// すでに登録されたトークン文字列を指定すると、実装によってはエラーを返す可能性があります。
	AddToken(ctx context.Context, token domain.Token) error
	// RemoveToken は指定されたトークン文字列を削除します。
	RemoveToken(ctx context.Context, token string) error
	// LoadTokens は現在登録されているトークンをすべて返します。
	LoadTokens(ctx context.Context) ([]*domain.Token, error)
	// ValidateToken は指定されたトークン文字列が登録されているか判定します。
	ValidateToken(ctx context.Context, token string) (bool, error)
}
