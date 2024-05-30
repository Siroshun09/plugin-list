package sqlite

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CustomDataKey(t *testing.T) {
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

	repo, err := c.NewCustomDataRepository(context.Background())
	assertion.Nil(err)

	// 初期状態では GetKeys は空の配列を返す
	keys, err := repo.GetKeys(context.Background())
	assertion.Nil(err)
	assertion.Equal(0, len(keys))

	// 初期状態では CustomDataKey の検索は nil を返す
	searchResult, err := repo.SearchForKey(context.Background(), "test")
	assertion.Nil(err)
	assertion.Nil(searchResult)

	// 初期状態ではキーの存在チェックは false を返す
	ok, err := repo.ExistsKey(context.Background(), "test")
	assertion.Nil(err)
	assertion.False(ok)

	key := domain.CustomDataKey{Key: "test", FormType: "text"} // DisplayName と Description がないキー

	assertion.Nil(repo.AddOrUpdateKey(context.Background(), key)) // 保存

	// GetKeys に含まれているか
	keys, err = repo.GetKeys(context.Background())
	assertion.Nil(err)
	assertion.Equal(1, len(keys))
	assertion.Equal(key, keys[0])

	// 検索でヒットするか
	searchResult, err = repo.SearchForKey(context.Background(), "test")
	assertion.Nil(err)
	assertion.Equal(key, *searchResult)

	// 存在チェックで true を返すか
	ok, err = repo.ExistsKey(context.Background(), "test")
	assertion.Nil(err)
	assertion.True(ok)

	key.DisplayName = "Test Key"
	key.Description = "This is a test key."

	assertion.Nil(repo.AddOrUpdateKey(context.Background(), key)) // 更新

	// GetKeys および検索の返り値が更新されているか
	keys, err = repo.GetKeys(context.Background())
	assertion.Nil(err)
	assertion.Equal(1, len(keys))
	assertion.Equal(key, keys[0])

	searchResult, err = repo.SearchForKey(context.Background(), "test")
	assertion.Nil(err)
	assertion.Equal(key, *searchResult)
}

func Test_PluginCustomData(t *testing.T) {
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

	repo, err := c.NewCustomDataRepository(context.Background())
	assertion.Nil(err)

	// 初期状態では GetPluginInfo は空の配列を返す
	data, err := repo.GetPluginInfo(context.Background(), "TestPlugin")
	assertion.Nil(err)
	assertion.Equal(0, len(data))

	descriptionKey := domain.CustomDataKey{Key: "description"}
	urlKey := domain.CustomDataKey{Key: "url"}

	// 事前にデータキーを保存しておく
	assertion.Nil(repo.AddOrUpdateKey(context.Background(), descriptionKey))
	assertion.Nil(repo.AddOrUpdateKey(context.Background(), urlKey))

	descriptionData := domain.PluginCustomData{Key: descriptionKey.Key, Data: "A plugin for Testing"}
	assertion.Nil(repo.AddOrUpdatePluginInfo(context.Background(), "TestPlugin", descriptionData)) // 保存

	// GetPluginInfo に含まれているか
	data, err = repo.GetPluginInfo(context.Background(), "TestPlugin")
	assertion.Nil(err)
	assertion.Equal(1, len(data))
	assertion.Equal(descriptionData, data[0])

	// 新しいデータを追加する
	urlData := domain.PluginCustomData{Key: urlKey.Key, Data: "https://example.com/TestPlugin"}
	assertion.Nil(repo.AddOrUpdatePluginInfo(context.Background(), "TestPlugin", urlData)) // 保存

	data, err = repo.GetPluginInfo(context.Background(), "TestPlugin")
	assertion.Nil(err)
	assertion.Equal(2, len(data))
	assertion.Equal(descriptionData, data[0])
	assertion.Equal(urlData, data[1])

	// 既存のデータを更新する
	urlData.Data = "https://example.com/plugins/TestPlugin"
	assertion.Nil(repo.AddOrUpdatePluginInfo(context.Background(), "TestPlugin", urlData)) // 更新

	data, err = repo.GetPluginInfo(context.Background(), "TestPlugin")
	assertion.Nil(err)
	assertion.Equal(2, len(data))
	assertion.Equal(descriptionData, data[0])
	assertion.Equal(urlData, data[1])
}
