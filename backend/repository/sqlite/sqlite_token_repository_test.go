package sqlite

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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

	repo, err := c.NewTokenRepository() // Create `tokens` table

	assertion.Nil(err)

	assertion.False(repo.ValidateToken(context.TODO(), "no tokens registered")) // At this time, no tokens are registered in the database

	token := domain.Token{Value: "First token", Created: time.UnixMilli(100)}
	assertion.Nil(repo.AddToken(context.TODO(), token))

	assertion.True(repo.ValidateToken(context.TODO(), token.Value))
	assertion.False(repo.ValidateToken(context.TODO(), "Invalid token"))

	tokens, err := repo.LoadTokens(context.TODO())

	assertion.Nil(err)
	assertion.Equal(1, len(tokens))
	assertion.Equal(&token, tokens[0])

	assertion.Nil(repo.RemoveToken(context.TODO(), token.Value))

	assertion.False(repo.ValidateToken(context.TODO(), token.Value))

	tokens, err = repo.LoadTokens(context.TODO())

	assertion.Nil(err)
	assertion.Equal(0, len(tokens))
}
