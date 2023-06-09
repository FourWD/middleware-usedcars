package orm

import orm "github.com/HinekoTech/middleware/orm"

type MenuGroup struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(45);" `

	Code string ` db:"code" json:"code" gorm:"type:varchar(45);" `
}
