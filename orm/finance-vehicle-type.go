package orm

import orm "github.com/HinekoTech/middleware/orm"

type FinanceVehicleType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DomainCode string ` db:"domain_code" json:"domain_code" gorm:"type:varchar(50);index " `

	FinanceID     string  ` db:"finance_id" json:"finance_id" gorm:"type:varchar(36); " `
	VehicleTypeID string  ` db:"vehicle_type_id" json:"vehicle_type_id" gorm:"type:varchar(36);  " `
	Month         int     ` db:"month" json:"month" gorm:"type:varchar(50);  " `
	Interest      float64 ` db:"interest" json:"interest" gorm:"type:float; " `
	RowOrder      int     ` db:"row_order" json:"row_order" gorm:"type:int(11); " `
}
