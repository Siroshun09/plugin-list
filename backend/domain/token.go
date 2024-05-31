package domain

import "time"

// Token は一部の API のアクセスに必要なトークンを保持します。
type Token struct {
	// 文字列でのトークン表記
	Value string
	// このトークンが作成された日時
	Created time.Time
}
