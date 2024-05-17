package sqlite

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestMCPluginRepository は SQLite データベースを使用した repository.TokenRepository の実装をテストします。
func TestTokenRepository(t *testing.T) {
	assertion := assert.New(t)
	d := t.TempDir()
	c, err := CreateConnection(d + "/sqlite.db")

	defer func(c Connection) {
		err := c.Close()
		if err != nil {
			t.Fatalf("failed to close database %s", err)
		}
	}(c)

	assertion.Nil(err)

	repo, err := c.NewTokenRepository(context.TODO()) // TokenRepository を作成 (データベースへのテーブル作成)

	assertion.Nil(err)

	// 初期段階では何のトークンも登録されていないため、ValidateToken は false を返すことが期待される
	assertion.False(repo.ValidateToken(context.TODO(), "no tokens registered"))

	// トークンの登録テスト
	token := domain.Token{Value: "First token", Created: time.UnixMilli(100)}
	assertion.Nil(repo.AddToken(context.TODO(), token))

	assertion.True(repo.ValidateToken(context.TODO(), token.Value))      // 登録したトークンであれば ValidateToken は true を返す
	assertion.False(repo.ValidateToken(context.TODO(), "Invalid token")) // 関係のない文字列は引き続き false を返す

	// トークンの一覧取得テスト
	// 返される配列は長さ1で、登録した "First token" が含まれる
	tokens, err := repo.LoadTokens(context.TODO())

	assertion.Nil(err)
	assertion.Equal(1, len(tokens))
	assertion.Equal(&token, tokens[0])

	// トークンの削除テスト
	assertion.Nil(repo.RemoveToken(context.TODO(), token.Value))

	assertion.False(repo.ValidateToken(context.TODO(), token.Value)) // 削除後、当該トークンに対する ValidateToken は false を返す

	// トークンの一覧取得テスト
	// 返される配列は空であることが期待される
	tokens, err = repo.LoadTokens(context.TODO())

	assertion.Nil(err)
	assertion.Equal(0, len(tokens))
}
