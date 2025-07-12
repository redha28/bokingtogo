package customerdto

type FamilyResponse struct {
	FlName     string `json:"fl_name"`
	FlDob      string `json:"fl_dob"`
	FlRelation string `json:"fl_relation"`
}

type CustomerData struct {
	CstID         int              `json:"cst_id"`
	CstName       string           `json:"cst_name"`
	CstEmail      string           `json:"cst_email"`
	CstDob        string           `json:"cst_dob"`
	CstPhoneNum   string           `json:"cst_phoneNum"`
	NationalityID int              `json:"nationality_id"`
	FamilyList    []FamilyResponse `json:"family_list"`
}

type CustomerResponse struct {
	Data CustomerData `json:"data"`
}

type CustomerListResponse struct {
	Data []CustomerData `json:"data"`
}
