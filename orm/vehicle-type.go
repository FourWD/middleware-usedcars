package orm

import orm "github.com/HinekoTech/middleware/orm"

type VehicleType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `db:"name" json:"type_name"`
	NameEn string `db:"name_en" json:"type_name_en"`
}
