package orm

import orm "github.com/HinekoTech/middleware/orm"

type PermissionMenu struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PermissionID string `db:"permission_id"  json:"permission_id" gorm:"type:varchar(36);index" `
	MenuID       string `db:"menu_id"  json:"menu_id" gorm:"type:varchar(36);index"`
	// ActionType   int       `db:"action_type" json:"action_type"  gorm:"default:null;type:tinyint(1)"`
}
