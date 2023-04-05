package orm

import orm "github.com/HinekoTech/middleware/orm"

type EngineOil struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name"`
	NameEng  string `db:"name_eng" json:"name_eng"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int(11);"`
}
