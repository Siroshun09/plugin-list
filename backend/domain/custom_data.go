package domain

// CustomDataKey は PluginCustomData に使用されるキーの詳細な情報を保持します
type CustomDataKey struct {
	// PluginCustomData の保存に使用されるキー
	Key string
	// 表示名
	DisplayName string
	// 説明
	Description string
	// 入力フォームの種類
	FormType string
}

// PluginCustomData はプラグインの任意の情報を保持します
type PluginCustomData struct {
	// データのキー
	Key string
	// 値
	Data string
}
