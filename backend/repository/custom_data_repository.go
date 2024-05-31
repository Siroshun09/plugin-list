package repository

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
)

// CustomDataRepository は domain.CustomDataKey および domain.PluginCustomData の管理を行います。
type CustomDataRepository interface {
	// GetKeys は現在登録されている domain.CustomDataKey をすべて返します
	GetKeys(ctx context.Context) ([]domain.CustomDataKey, error)
	// SearchForKey は引数に一致するキーを返します。見つからない場合は nil が返されます
	SearchForKey(ctx context.Context, key string) (*domain.CustomDataKey, error)
	// AddOrUpdateKey は指定したキーを追加または更新します
	AddOrUpdateKey(ctx context.Context, key domain.CustomDataKey) error
	// ExistsKey は指定したキーが存在するかを判定します
	ExistsKey(ctx context.Context, key string) (bool, error)
	// GetPluginInfo は指定したプラグインの domain.PluginCustomData をすべて返します
	GetPluginInfo(ctx context.Context, pluginName string) ([]domain.PluginCustomData, error)
	// AddOrUpdatePluginInfo は指定したプラグインに domain.PluginCustomData を追加または更新します
	AddOrUpdatePluginInfo(ctx context.Context, pluginName string, data domain.PluginCustomData) error
}
