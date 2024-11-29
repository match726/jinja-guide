package models

import (
	"context"
	"fmt"
)

type ShrinesListResp struct {
	Name            string   `json:"name"`
	Address         string   `json:"address"`
	PlusCode        string   `json:"plusCode"`
	PlaceID         string   `json:"placeId"`
	ObjectOfWorship []string `json:"objectOfWorship"`
	HasGoshuin      bool     `json:"hasGoshuin"`
}

func (pg *Postgres) GetShrinesListByStdAreaCode(sacr *SacRelationship) (shrlrs []*ShrinesListResp, err error) {

	var query string

	switch sacr.Kinds {
	case "Pref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id, shrc.content1
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.pref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "SubPref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id, shrc.content1
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.subpref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "City", "District":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id, shrc.content1
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code1 = $1
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "Town/Village", "Ward":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id, shrc.content1
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code2 = $1
						AND shr.std_area_code = sac.std_area_code
					LEFT JOIN t_shrine_contents shrc
						ON shrc.id = 8
						AND shr.plus_code = shrc.keyword1
					ORDER BY shr.std_area_code, shr.address, shr.name`
	}

	rows, err := pg.dbPool.Query(context.Background(), query, sacr.StdAreaCode)
	if err != nil {
		return nil, fmt.Errorf("神社一覧 取得失敗： %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var shrlr ShrinesListResp

		err = rows.Scan(&shrlr.Name, &shrlr.Address, &shrlr.PlusCode, &shrlr.PlaceID, &shrlr.ObjectOfWorship)
		if err != nil {
			return nil, fmt.Errorf("スキャン失敗： %w", err)
		}

		shrlrs = append(shrlrs, &shrlr)

	}

	return shrlrs, err

}
