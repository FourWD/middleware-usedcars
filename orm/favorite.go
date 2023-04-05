package orm

import orm "github.com/HinekoTech/middleware/orm"

type Favorite struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DomainCode string `db:"domain_code"  json:"domain_code" gorm:"index;default:null; type:varchar(50)"`

	Code     string `db:"code"  json:"code" gorm:"index; default:null; type:varchar(50)"`
	UserID   string `db:"user_id" json:"user_id" gorm:"type:varchar(36);"`
	ObjectID string `db:"object_id" json:"object_id" gorm:"type:varchar(36);"`
}
