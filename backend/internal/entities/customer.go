package entities

type Customer struct {
	CstID         int      `gorm:"primaryKey;column:cst_id" json:"cst_id"`
	CstName       string   `gorm:"column:cst_name" json:"cst_name"`
	CstEmail      string   `gorm:"column:cst_email" json:"cst_email"`
	CstDob        string   `gorm:"column:cst_dob" json:"cst_dob"`
	CstPhoneNum   string   `gorm:"column:cst_phone_num" json:"cst_phoneNum"`
	NationalityID int      `gorm:"column:nationality_id" json:"nationality_id"`
	FamilyList    []Family `gorm:"foreignKey:CstID" json:"family_list"`
}
