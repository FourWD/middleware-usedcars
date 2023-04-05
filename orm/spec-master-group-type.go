package orm

import orm "github.com/HinekoTech/middleware/orm"

type SpecMasterGroupType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `db:"name"  json:"name" gorm:"default:null;type:varchar(200)"`
	NameEn string `db:"name_en"  json:"name_en" gorm:"default:null;type:varchar(200)"`
	// feature & spec
}
