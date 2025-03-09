package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/match726/jinja-guide/tree/main/server/domain/model"
)

func SendMessage(url string, token string, content string) error {

	// リクエストボディの設定
	reqBody := model.ShrineRegisterErrMessage{
		Content: content,
	}
	reqBodyJson, _ := json.Marshal(reqBody)

	// リクエスト作成
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bytes.NewBuffer(reqBodyJson),
	)
	if err != nil {
		return fmt.Errorf("[リクエスト作成失敗]: %w", err)
	}

	// リクエストヘッダー設定
	req.Header.Add("Authorization", "Bot "+token)
	req.Header.Add("Content-Type", "application/json")

	// POST実行
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[POSTリクエスト失敗]: %w", err)
	}
	defer resp.Body.Close()

	return nil

}
