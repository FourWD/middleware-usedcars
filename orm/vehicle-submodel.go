package orm

import orm "github.com/HinekoTech/middleware/orm"

type VehicleSubmodel struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name           string `db:"name"  json:"name" gorm:"index;not null;type:varchar(50)"`
	VehicleModelID string `db:"vehicle_model_id"  json:"vehicle_model_id" gorm:"type:varchar(36)"`
}
