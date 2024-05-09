package api

import (
	"bytes"
	"encoding/json"
	"github.com/Siroshun09/plugin-list/domain"
	mockUsecase "github.com/Siroshun09/plugin-list/usecase/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestConvertMCPluginAndPlugin は domain.MCPlugin と api.Plugin の変換をテストします
func TestConvertMCPluginAndPlugin(t *testing.T) {
	assertions := assert.New(t)

	mcPlugin := createTestMCPlugin()
	plugin := createTestPlugin()

	assertions.Equal(*mcPlugin, toMCPlugin(plugin)) // Test converted MCPlugin from Plugin
	assertions.Equal(*plugin, toPlugin(mcPlugin))   // Test converted Plugin from MCPlugin
}

// TestGetPluginsByServer は PluginList.GetPluginsByServer をテストします
func TestGetPluginsByServer(t *testing.T) {
	assertions := assert.New(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/servers/test/plugins", nil) // /servers/test/plugins にリクエストが来たと想定します

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockUsecase.NewMockMCPluginUseCase(ctrl)
	plList := NewPluginList(m)

	mcPlugins := make([]*domain.MCPlugin, 1)
	mcPlugins[0] = createTestMCPlugin()
	m.EXPECT().GetMCPluginsByServerName(gomock.Any(), "test").Return(mcPlugins, nil) // MCPluginUseCase.GetMCPluginsByServerName はサーバー名 test で呼び出されることが期待されます

	plList.GetPluginsByServer(w, r, "test") // サーバー名 test で当該メソッドを呼び出します

	resp := w.Result()
	defer assertions.NoError(resp.Body.Close())
	assertions.Equal(http.StatusOK, resp.StatusCode) // 通常、ステータスコード 200 を返します

	// 返されたプラグインの一覧を確認します。
	// モックによって1つの TestPlugin を含んだ配列を返すようにしたので
	// レスポンスのボディの JSON からも同じ内容の配列が返されることを期待します。
	var result []*Plugin
	assertions.Nil(json.NewDecoder(resp.Body).Decode(&result))
	assertions.Equal(1, len(result))
	assertions.Equal(*mcPlugins[0], toMCPlugin(result[0]))
}

func TestAddPlugins(t *testing.T) {
	assertions := assert.New(t)
	plugin := createTestPlugin()
	body := "[{\"plugin_name\":\"TestPlugin\",\"file_name\":\"TestPlugin-1.0.0.jar\",\"last_updated\":100,\"type\":\"bukkit_plugin\",\"version\":\"1.0.0\"}]"

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/servers/test/plugins", bytes.NewBuffer([]byte(body))) // /servers/test/plugins にリクエストが来たと想定します

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockUsecase.NewMockMCPluginUseCase(ctrl)
	plList := NewPluginList(m)
	mcPlugin := toMCPlugin(plugin)

	m.EXPECT().SubmitMCPlugin(gomock.Any(), &mcPlugin).Return(nil) // MCPluginUseCase.SubmitMCPlugin が TestPlugin を引数として呼び出されることを期待します

	plList.AddPlugins(w, r, plugin.ServerName)

	resp := w.Result()

	defer assertions.NoError(resp.Body.Close())
	assertions.Equal(http.StatusCreated, resp.StatusCode) // 通常、ステータスコード 201 を返します
}

// TestAddPlugin は PluginList.AddPlugin をテストします
func TestAddPlugin(t *testing.T) {
	assertions := assert.New(t)
	plugin := createTestPlugin()
	body := "{\"file_name\":\"TestPlugin-1.0.0.jar\",\"last_updated\":100,\"type\":\"bukkit_plugin\",\"version\":\"1.0.0\"}"

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/servers/test/plugins/TestPlugin", bytes.NewBuffer([]byte(body))) // /servers/test/plugins/TestPlugin にリクエストが来たと想定します

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockUsecase.NewMockMCPluginUseCase(ctrl)
	plList := NewPluginList(m)
	mcPlugin := toMCPlugin(plugin)

	m.EXPECT().SubmitMCPlugin(gomock.Any(), &mcPlugin).Return(nil) // MCPluginUseCase.SubmitMCPlugin が TestPlugin を引数として呼び出されることを期待します

	plList.AddPlugin(w, r, plugin.ServerName, plugin.PluginName)

	resp := w.Result()

	defer assertions.NoError(resp.Body.Close())
	assertions.Equal(http.StatusCreated, resp.StatusCode) // 通常、ステータスコード 201 を返します
}

// TestRemovePlugin は PluginList.RemovePlugin をテストします
func TestRemovePlugin(t *testing.T) {
	assertions := assert.New(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/servers/test/plugins/TestPlugin", nil) // /servers/test/plugins/TestPlugin にリクエストが来たと想定します

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockUsecase.NewMockMCPluginUseCase(ctrl)
	plList := NewPluginList(m)

	// MCPluginUseCase.DeleteMCPlugin はサーバー名 test, プラグイン名 TestPlugin で呼び出されることが期待されます
	m.EXPECT().DeleteMCPlugin(gomock.Any(), "test", "TestPlugin").Return(nil)

	plList.DeletePlugin(w, r, "test", "TestPlugin") // サーバー名 test, プラグイン名 TestPlugin で当該メソッドを呼び出します

	resp := w.Result()
	assertions.Equal(http.StatusNoContent, resp.StatusCode) // 通常、ステータスコードは 204 を返します
}

func TestGetServerNames(t *testing.T) {
	assertions := assert.New(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/servers", nil) // /servers にリクエストが来たと想定します

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mockUsecase.NewMockMCPluginUseCase(ctrl)
	plList := NewPluginList(m)

	serverNames := []string{"test_1", "test_2"}
	m.EXPECT().GetServerNames(gomock.Any()).Return(serverNames, nil)

	plList.GetServerNames(w, r)

	resp := w.Result()

	defer assertions.NoError(resp.Body.Close())
	assertions.Equal(http.StatusOK, resp.StatusCode) // 通常、ステータスコード 200 を返します

	var result []string
	assertions.NoError(json.NewDecoder(resp.Body).Decode(&result))
	assertions.Equal(serverNames, result)
}

func createTestMCPlugin() *domain.MCPlugin {
	return &domain.MCPlugin{
		PluginName:  "TestPlugin",
		FileName:    "TestPlugin-1.0.0.jar",
		Version:     "1.0.0",
		Type:        "bukkit_plugin",
		ServerName:  "test",
		LastUpdated: time.UnixMilli(100),
	}
}

func createTestPlugin() *Plugin {
	return &Plugin{
		PluginName:  "TestPlugin",
		FileName:    "TestPlugin-1.0.0.jar",
		Version:     "1.0.0",
		Type:        "bukkit_plugin",
		ServerName:  "test",
		LastUpdated: 100,
	}
}
