package orm

import orm "github.com/HinekoTech/middleware/orm"

type Permission struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name string `db:"name"  json:"name" gorm:"dafault:null;type:varchar(45)"`
}
