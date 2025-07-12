package entities

type Nationality struct {
	NationalityID   int    `gorm:"primaryKey;column:nationality_id" json:"nationality_id"`
	NationalityName string `gorm:"column:nationality_name" json:"nationality_name"`
	NationalityCode string `gorm:"column:nationality_code" json:"nationality_code"`
}
