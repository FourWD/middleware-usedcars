package orm

import orm "github.com/HinekoTech/middleware/orm"

type Insurance struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(100);index" `

	Name     string ` db:"name" json:"name" gorm:"type:varchar(45);" `
	RowOrder int    ` db:"row_order" json:"row_order" gorm:"type:int(11);" `
}
