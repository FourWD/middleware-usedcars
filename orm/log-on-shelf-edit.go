package orm

import orm "github.com/HinekoTech/middleware/orm"

type LogOnShelfEdit struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	OnShelf bool `db:"on_shelf"  json:"on_shelf" gorm:"type:tinyint(1)"`
}
