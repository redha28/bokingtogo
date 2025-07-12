package familydto

type CreateFamilyRequest struct {
	CstID      int    `json:"cst_id" validate:"required"`
	FlName     string `json:"fl_name" validate:"required"`
	FlDob      string `json:"fl_dob" validate:"required"`
	FlRelation string `json:"fl_relation" validate:"required"`
}

type UpdateFamilyRequest struct {
	FlName     string `json:"fl_name" validate:"required"`
	FlDob      string `json:"fl_dob" validate:"required"`
	FlRelation string `json:"fl_relation" validate:"required"`
}
