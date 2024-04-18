package usecase

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
)

// MCPluginUseCase はプラグイン情報の取得や更新を行います。
type MCPluginUseCase interface {
	// GetMCPluginsByServerName はサーバー名を指定して、そのサーバーに導入されているプラグインの配列を取得します。
	GetMCPluginsByServerName(ctx context.Context, serverName string) ([]*domain.MCPlugin, error)
	// DeleteMCPlugin  は指定されたプラグイン名・サーバー名に紐づけられた MCPlugin の情報を削除します。
	DeleteMCPlugin(ctx context.Context, serverName string, pluginName string) error
	// SubmitMCPlugin はプラグイン情報を作成 (Create) または更新 (Update) します。
	SubmitMCPlugin(ctx context.Context, plugin *domain.MCPlugin) error
}
