package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type WikiMediaResp struct {
	OriginalImage struct {
		Source string `json:"source"`
	} `json:"originalimage"`
	Extarct string `json:"extract"`
}

var wikiMediaResp WikiMediaResp

func GetShrineDetailsFromWikipedia(url string) (image string, extract string, err error) {

	title := url[strings.LastIndex(url, "/")+1:]
	resp, err := http.Get("https://ja.wikipedia.org/api/rest_v1/page/summary/" + title)
	if err != nil {
		return "", "", fmt.Errorf("APIリクエスト失敗: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("Body取得失敗: %w", err)
	}
	defer resp.Body.Close()

	json.Unmarshal(body, &wikiMediaResp)

	return wikiMediaResp.OriginalImage.Source, wikiMediaResp.Extarct, nil

}
