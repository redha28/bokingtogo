package nationalitydto

type CreateNationalityRequest struct {
	NationalityName string `json:"nationality_name" validate:"required"`
	NationalityCode string `json:"nationality_code" validate:"required,len=2"`
}

type UpdateNationalityRequest struct {
	NationalityName string `json:"nationality_name" validate:"required"`
	NationalityCode string `json:"nationality_code" validate:"required,len=2"`
}
