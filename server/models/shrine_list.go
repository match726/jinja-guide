package models

import (
	"context"
	"fmt"
)

func (pg *Postgres) GetShrinesListByStdAreaCode(sacr *SacRelationship) (shrs []*Shrine, err error) {

	var query string

	switch sacr.Kinds {
	case "Pref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.pref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "SubPref":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.subpref_area_code = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "City", "District":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code1 = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	case "Town/Village", "Ward":
		query = `SELECT shr.name, shr.address, shr.plus_code, shr.place_id
					FROM t_shrines shr
					INNER JOIN m_stdareacode sac
						ON sac.munic_area_code2 = $1
						AND shr.std_area_code = sac.std_area_code
					ORDER BY shr.std_area_code, shr.address, shr.name`
	}

	rows, err := pg.dbPool.Query(context.Background(), query, sacr.StdAreaCode)
	if err != nil {
		return nil, fmt.Errorf("神社一覧 取得失敗： %w", err)
	}
	defer rows.Close()

	for rows.Next() {

		var shr Shrine

		err = rows.Scan(&shr.Name, &shr.Address, &shr.PlusCode, &shr.PlaceID)
		if err != nil {
			return nil, fmt.Errorf("スキャン失敗： %w", err)
		}

		shrs = append(shrs, &shr)

	}

	return shrs, err

}
