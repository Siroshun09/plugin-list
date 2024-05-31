package domain

import "time"

// MCPlugin はプラグインの情報を表します。
type MCPlugin struct {
	// プラグインの名前
	PluginName string
	// このプラグインが導入されているサーバーの名前
	ServerName string
	// ファイル名
	FileName string
	// バージョン
	Version string
	// プラグインの種類
	// 例: `bukkit_plugin`, `paper_plugin`, `velocity_plugin`
	Type string
	// インストールしたプラグインの更新日
	LastUpdated time.Time
}
