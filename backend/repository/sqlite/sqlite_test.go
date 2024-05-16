package sqlite

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// TestSqliteConnection は SQLite データベースファイルを作成し、接続を確立するまでをテストします。
func TestSqliteConnection(t *testing.T) {
	assertion := assert.New(t)
	d := t.TempDir()
	dbPath := d + "/" + DatabaseFilename

	connectToDatabaseAndCheckConnection(t, assertion, dbPath) // 初回接続テスト

	// データベースファイルが存在するか確認
	_, err := os.Stat(dbPath)
	assertion.Nil(err)

	connectToDatabaseAndCheckConnection(t, assertion, dbPath) // 既存のデータベースファイルへの接続テスト
}

func connectToDatabaseAndCheckConnection(t *testing.T, assertion *assert.Assertions, filepath string) {
	c, err := CreateConnection(filepath)

	defer func(c Connection) {
		err := c.Close()
		if err != nil {
			t.Fatalf("failed to close database %s", err)
		}
	}(c)

	assertion.Nil(err)
	assertion.True(c.IsOpen())
}
