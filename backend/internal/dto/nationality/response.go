package nationalitydto

type NationalityData struct {
	NationalityID   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}

type NationalityResponse struct {
	Data NationalityData `json:"data"`
}

type NationalityListResponse struct {
	Data []NationalityData `json:"data"`
}
