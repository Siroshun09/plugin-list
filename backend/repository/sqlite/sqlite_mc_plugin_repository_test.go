package sqlite

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestMCPluginRepository は SQLite データベースを使用した repository.MCPluginRepository の実装をテストします。
func TestMCPluginRepository(t *testing.T) {
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

	repo, err := c.NewMCPluginRepository(context.TODO()) // MCPluginRepository を作成 (データベースへのテーブル作成)
	assertion.Nil(err)

	// このテストで使用するサンプルのプラグイン情報
	testPlugin := domain.MCPlugin{
		PluginName:  "TestPlugin",
		FileName:    "TestPlugin-1.0.jar",
		Version:     "1.0",
		Type:        "bukkit_plugin",
		ServerName:  "test",
		LastUpdated: time.UnixMilli(100),
	}

	// プラグイン情報の保存テスト
	err = repo.CreateOrUpdateMCPlugin(context.TODO(), testPlugin)
	assertion.Nil(err)
	checkTestPlugin(repo, &testPlugin, assertion) // 保存したプラグイン情報とサンプルが同じか検証

	// プラグイン情報の更新テスト
	testPlugin.FileName = "TestPlugin-1.1.jar"
	testPlugin.Version = "1.1"
	testPlugin.LastUpdated = time.UnixMilli(300)

	err = repo.CreateOrUpdateMCPlugin(context.TODO(), testPlugin)

	assertion.Nil(err)
	checkTestPlugin(repo, &testPlugin, assertion) // 更新したプラグイン情報とサンプルが同じか検証

	// サーバーの名前一覧の取得テスト
	// 返される配列は長さ1で、"test" が含まれる
	serverNames, err := repo.GetServerNames(context.TODO())

	assertion.Nil(err)
	assertion.Equal(1, len(serverNames))
	assertion.Equal(testPlugin.ServerName, serverNames[0])

	// プラグイン情報の削除テスト
	err = repo.DeleteMCPlugin(context.TODO(), testPlugin.ServerName, testPlugin.PluginName)

	assertion.Nil(err)

	// 削除後のプラグイン一覧の取得テスト
	// 返される配列は空であることが期待される
	plugins, err := repo.GetMCPluginsByServerName(context.TODO(), testPlugin.ServerName)

	assertion.Nil(err)
	assertion.Equal(0, len(plugins))
}

// checkTestPlugin はリポジトリから取得した MCPlugin と与えられた MCPlugin が等しいかどうかを確認するヘルパーメソッドです。
// このメソッドを呼び出す際、リポジトリには1つの MCPlugin のみが含まれている必要があります。
func checkTestPlugin(repo repository.MCPluginRepository, testPlugin *domain.MCPlugin, assertion *assert.Assertions) {
	plugins, err := repo.GetMCPluginsByServerName(context.TODO(), testPlugin.ServerName) // リポジトリからプラグイン一覧を取得

	assertion.Nil(err)

	// 返されるプラグインの配列は長さ1で、引数のプラグインと同等のものが含まれる
	assertion.Equal(1, len(plugins))
	assertion.Equal(testPlugin, plugins[0])
}
