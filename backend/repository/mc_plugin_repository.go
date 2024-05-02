package repository

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
)

type MCPluginRepository interface {
	// CreateOrUpdateMCPlugin はプラグイン情報を作成または更新します。
	CreateOrUpdateMCPlugin(ctx context.Context, plugin *domain.MCPlugin) error
	// DeleteMCPlugin  は指定されたプラグイン名・サーバー名に紐づけられた MCPlugin の情報を削除します。
	DeleteMCPlugin(ctx context.Context, serverName string, pluginName string) error
	// GetMCPluginsByServerName はサーバー名を指定して、そのサーバーに導入されているプラグインの配列を取得します。
	GetMCPluginsByServerName(ctx context.Context, serverName string) ([]*domain.MCPlugin, error)
	// GetServerNames は記録されているプラグインのサーバー名をすべて取得します。
	GetServerNames(ctx context.Context) ([]string, error)
}
