package customerdto

type FamilyRequest struct {
	FlName     string `json:"fl_name" validate:"required"`
	FlDob      string `json:"fl_dob" validate:"required"`
	FlRelation string `json:"fl_relation" validate:"required"`
}

type CreateCustomerRequest struct {
	CstName       string          `json:"cst_name" validate:"required"`
	CstEmail      string          `json:"cst_email" validate:"required,email"`
	CstDob        string          `json:"cst_dob" validate:"required"`
	CstPhoneNum   string          `json:"cst_phoneNum" validate:"required"`
	NationalityID int             `json:"nationality_id" validate:"required"`
	FamilyList    []FamilyRequest `json:"family_list"`
}
