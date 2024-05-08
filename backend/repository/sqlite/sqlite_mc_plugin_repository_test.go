package sqlite

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMcPluginRepository(t *testing.T) {
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

	repo, err := c.NewMCPluginRepository() // Create `mc_plugins` table

	assertion.Nil(err)

	testPlugin := domain.MCPlugin{
		PluginName:  "TestPlugin",
		FileName:    "TestPlugin-1.0.0.jar",
		Version:     "1.0.0",
		Type:        "bukkit_plugin",
		ServerName:  "test",
		LastUpdated: time.UnixMilli(100),
	}

	// Save `TestPlugin` data (New plugin)

	err = repo.CreateOrUpdateMCPlugin(context.TODO(), &testPlugin)

	assertion.Nil(err)
	checkTestPlugin(repo, &testPlugin, assertion)

	// Update `TestPlugin` data (Existing plugin)

	testPlugin.FileName = "TestPlugin-1.0.1.jar"
	testPlugin.Version = "1.0.1"

	err = repo.CreateOrUpdateMCPlugin(context.TODO(), &testPlugin)

	assertion.Nil(err)
	checkTestPlugin(repo, &testPlugin, assertion)

	// Get server names

	serverNames, err := repo.GetServerNames(context.TODO())

	assertion.Nil(err)
	assertion.Equal(1, len(serverNames))
	assertion.Equal(testPlugin.ServerName, serverNames[0])

	// Delete `TestPlugin` data

	err = repo.DeleteMCPlugin(context.TODO(), testPlugin.ServerName, testPlugin.PluginName)

	assertion.Nil(err)

	plugins, err := repo.GetMCPluginsByServerName(context.TODO(), testPlugin.ServerName)

	assertion.Nil(err)
	assertion.Equal(0, len(plugins))

}

// checkTestPlugin はリポジトリから取得した MCPlugin と与えられた MCPlugin が等しいかどうかを確認するヘルパーメソッドです。
// このメソッドを呼び出す際、リポジトリには1つの MCPlugin のみが含まれている必要があります。
func checkTestPlugin(repo repository.MCPluginRepository, testPlugin *domain.MCPlugin, assertion *assert.Assertions) {
	plugins, err := repo.GetMCPluginsByServerName(context.TODO(), testPlugin.ServerName) // Get the plugin list that include `TestPlugin`

	assertion.Nil(err)
	// The returned list has only `TestPlugin`
	assertion.Equal(1, len(plugins))
	assertion.Equal(testPlugin, plugins[0])
}
