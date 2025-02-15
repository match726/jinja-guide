package model

// Discord通知用構造体（神社登録エラー時）
type ShrineRegisterErrMessage struct {
	CreatedAt    string `json:"created_at"`
	ErrorDetails string `json:"error_details"`
	ShrineInfo   struct {
		Name          string `json:"name"`
		Address       string `json:"address"`
		PlaceID       string `json:"place_id"`
		GoogleMapLink string `json:"google_map_link"`
	} `json:"shrine_info"`
}
