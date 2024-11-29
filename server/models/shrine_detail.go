package models

import (
	"context"
	"fmt"
	"strconv"
)

// 神社の詳細情報を取得
func (pg *Postgres) GetShrineDetails(shr *Shrine) (shrd ShrineDetails, err error) {

	query1 := `SELECT shr.name, shr.address, shr.place_id
						FROM t_shrines shr
						WHERE shr.plus_code = $1`

	row := pg.dbPool.QueryRow(context.Background(), query1, shr.PlusCode)

	err = row.Scan(&shrd.Name, &shrd.Address, &shrd.PlaceID)
	if err != nil {
		return shrd, fmt.Errorf("スキャン１失敗： %w", err)
	}

	query2 := `SELECT shrc.id, shrc.content1, shrc.content2, shrc.content3
              FROM t_shrine_contents shrc
              WHERE shrc.keyword1 = $1
              ORDER BY shrc.id, shrc.keyword1, shrc.keyword2`

	rows, err := pg.dbPool.Query(context.Background(), query2, shr.PlusCode)
	if err != nil {
		return shrd, fmt.Errorf("神社詳細情報 取得失敗： %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var shrc ShrineContents

		err = rows.Scan(&shrc.Id, &shrc.Content1, &shrc.Content2, &shrc.Content3)
		if err != nil {
			return shrd, fmt.Errorf("スキャン失敗： %w", err)
		}

		switch shrc.Id {
		case 1:
			// 振り仮名の設定
			shrd.Furigana = shrc.Content1
		case 2:
			// 別名称の設定
			shrd.AltName = append(shrd.AltName, shrc.Content1)
		case 3:
			// 説明の設定
			shrd.Description = shrc.Content1
		case 4:
			// 関連タグの設定
			shrd.Tags = append(shrd.Tags, shrc.Content1)
		case 5:
			// 創建年の設定
			if _, err = strconv.Atoi(shrc.Content1); err == nil {
				shrd.FoundedYear = shrc.Content1 + "年"
			} else {
				shrd.FoundedYear = shrc.Content1
			}
		case 6:
			// 御祭神の設定
			shrd.ObjectOfWorship = append(shrd.ObjectOfWorship, shrc.Content1)
		case 7:
			// 社格の設定
			shrd.ShrineRank = append(shrd.ShrineRank, shrc.Content1)
		case 8:
			//御朱印の設定
			if shrc.Content1 == "あり" {
				shrd.HasGoshuin = true
			}
		case 9:
			// 公式サイトの設定
			shrd.WebsiteURL = shrc.Content1
		case 10:
			// Wikipediaの設定
			shrd.WikipediaURL = shrc.Content1
		}

	}

	// Wikipediaから情報取得
	if len(shrd.WikipediaURL) != 0 {
		image, extract, err := GetShrineDetailsFromWikipedia(shrd.WikipediaURL)
		if err != nil {
			return shrd, fmt.Errorf("%w", err)
		} else {
			shrd.Image = image
			if len(shrd.Description) == 0 {
				shrd.Description = extract
			}
		}
	}

	if len(shrd.AltName) == 0 {
		shrd.AltName = []string{"登録なし"}
	}
	if len(shrd.Tags) == 0 {
		shrd.Tags = []string{"登録なし"}
	}
	if len(shrd.ObjectOfWorship) == 0 {
		shrd.ObjectOfWorship = []string{"登録なし"}
	}
	if len(shrd.ShrineRank) == 0 {
		shrd.ShrineRank = []string{"登録なし"}
	}

	return shrd, err

}
