package sqlite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqliteConnection_IsOpen(t *testing.T) {
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
	assertion.True(c.IsOpen())
}
