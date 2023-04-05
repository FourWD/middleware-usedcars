package orm

import orm "github.com/HinekoTech/middleware/orm"

type Province struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `db:"name"  json:"name" gorm:"not null;index;type:varchar(150)"`
	NameEn   string `db:"name_en"  json:"name_en" gorm:"not null;index;type:varchar(150)"`
	RegionID string `db:"region_id"  json:"region_id" gorm:"index;type:varchar(36)"`
}
