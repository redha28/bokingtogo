package entities

type Family struct {
	FlID       int    `gorm:"primaryKey;column:fl_id" json:"fl_id"`
	CstID      int    `gorm:"column:cst_id" json:"cst_id"`
	FlName     string `gorm:"column:fl_name" json:"fl_name"`
	FlDob      string `gorm:"column:fl_dob" json:"fl_dob"`
	FlRelation string `gorm:"column:fl_relation" json:"fl_relation"`
}
