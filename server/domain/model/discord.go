package model

// Discord通知用構造体（神社登録エラー時）
type ShrineRegisterErrMessage struct {
	Content string `json:"content"`
}
