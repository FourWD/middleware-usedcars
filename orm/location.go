package orm

import orm "github.com/HinekoTech/middleware/orm"

type Location struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	LocationGroupID string `db:"location_group_id" json:"location_group_id" gorm:"type:varchar(36);" `
	Name            string ` db:"name" json:"name" gorm:"type:varchar(150);" `
}
