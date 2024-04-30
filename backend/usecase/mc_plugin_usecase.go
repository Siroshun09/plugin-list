package usecase

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
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

func NewMCPluginUseCase(repo repository.MCPluginRepository) MCPluginUseCase {
	return repositoryUseCase{repo}
}

type repositoryUseCase struct {
	repo repository.MCPluginRepository
}

func (r repositoryUseCase) GetMCPluginsByServerName(ctx context.Context, serverName string) ([]*domain.MCPlugin, error) {
	return r.repo.GetMCPluginsByServerName(ctx, serverName)
}

func (r repositoryUseCase) DeleteMCPlugin(ctx context.Context, serverName string, pluginName string) error {
	return r.repo.DeleteMCPlugin(ctx, serverName, pluginName)
}

func (r repositoryUseCase) SubmitMCPlugin(ctx context.Context, plugin *domain.MCPlugin) error {
	return r.repo.CreateOrUpdateMCPlugin(ctx, plugin)
}