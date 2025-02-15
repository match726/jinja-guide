package discord

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendMessage(url string, token string, content string) error {

	// メッセージ本文の設定
	contentJson := fmt.Sprintf(`{"content":"%s"}`, content)

	// リクエスト作成
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(contentJson)),
	)
	if err != nil {
		return fmt.Errorf("[リクエスト作成失敗]: %w", err)
	}

	// Content-Type 設定
	req.Header.Set("Authorization", "Bot "+token)
	req.Header.Set("Content-Type", "application/json")

	// POST実行
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[POSTリクエスト失敗]: %w", err)
	}
	defer resp.Body.Close()

	return nil

}
