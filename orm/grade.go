package orm

import orm "github.com/HinekoTech/middleware/orm"

type Grade struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(50);index " `

	Name     string ` db:"name" json:"name" gorm:"type:varchar(25); index " `
	Img      string ` db:"img" json:"img" gorm:"type:varchar(255);" `
	RowOrder int    ` db:"row_order" json:"row_order" gorm:"type:int(11);" `
}
