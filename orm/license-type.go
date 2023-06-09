package orm

import orm "github.com/HinekoTech/middleware/orm"

type LicenseType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `db:"name" json:"name" gorm:"type:varchar(45);"`
	NameEn string `db:"name_en" json:"name_en" gorm:"type:varchar(45);"`
}
