package familydto

type FamilyData struct {
	FlID       int    `json:"fl_id"`
	CstID      int    `json:"cst_id"`
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
	FlRelation string `json:"fl_relation"`
}

type FamilyResponse struct {
	Data FamilyData `json:"data"`
}

type FamilyListResponse struct {
	Data []FamilyData `json:"data"`
}
