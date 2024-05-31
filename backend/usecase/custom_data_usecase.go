package usecase

import (
	"context"
	"github.com/Siroshun09/plugin-list/domain"
	"github.com/Siroshun09/plugin-list/repository"
)

// CustomDataUseCase はプラグインのカスタムデータの取得や更新を行います。
type CustomDataUseCase interface {
	// GetKeys は現在登録されている domain.CustomDataKey をすべて返します
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
	AddOrUpdatePluginInfo(ctx context.Context, pluginName string, data []domain.PluginCustomData) error
}

// NewCustomDataUseCase は repository.CustomDataRepository を使用した新しい CustomDataUseCase を作成します。
func NewCustomDataUseCase(repo repository.CustomDataRepository) CustomDataUseCase {
	return customDataRepositoryUseCase{repo}
}

type customDataRepositoryUseCase struct {
	repo repository.CustomDataRepository
}

func (c customDataRepositoryUseCase) GetKeys(ctx context.Context) ([]domain.CustomDataKey, error) {
	return c.repo.GetKeys(ctx)
}

func (c customDataRepositoryUseCase) SearchForKey(ctx context.Context, key string) (*domain.CustomDataKey, error) {
	return c.repo.SearchForKey(ctx, key)
}

func (c customDataRepositoryUseCase) AddOrUpdateKey(ctx context.Context, key domain.CustomDataKey) error {
	return c.repo.AddOrUpdateKey(ctx, key)
}

func (c customDataRepositoryUseCase) ExistsKey(ctx context.Context, key string) (bool, error) {
	return c.repo.ExistsKey(ctx, key)
}

func (c customDataRepositoryUseCase) GetPluginInfo(ctx context.Context, pluginName string) ([]domain.PluginCustomData, error) {
	return c.repo.GetPluginInfo(ctx, pluginName)
}

func (c customDataRepositoryUseCase) AddOrUpdatePluginInfo(ctx context.Context, pluginName string, data []domain.PluginCustomData) error {
	for _, customData := range data {
		if err := c.repo.AddOrUpdatePluginInfo(ctx, pluginName, customData); err != nil {
			return err
		}
	}
	return nil
}
