package product_entity

import (
	"database/sql"
	"go-master-data/entity"
)

type ProductEntity struct {
	entity.AbstractEntity
	CompanyID    sql.NullInt64
	CategoryID   sql.NullInt64
	GroupID      sql.NullInt64
	Code         sql.NullString
	Name         sql.NullString
	SellingPrice sql.NullFloat64
	BuyingPrice  sql.NullFloat64
	Uom1         sql.NullString `gorm:"column:uom_1"`
	Uom2         sql.NullString `gorm:"column:uom_2"`
	Conv1To2     sql.NullInt32  `gorm:"column:conv_1_to_2"`
}

type ProductDetailEntity struct {
	ProductEntity
	CategoryCode sql.NullString
	CategoryName sql.NullString
	GroupCode    sql.NullString
	GroupName    sql.NullString
}

func (ProductEntity) TableName() string {
	return "product"
}
