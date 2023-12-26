package product_entity

import "go-master-data/entity"

type ProductCategoryEntity struct {
	entity.AbstractEntity
	Code       string
	Name       string
	CompanyID  int64
	DivisionID int64
}

type ProductCategoryDetailEntity struct {
	ProductCategoryEntity
}

func (ProductCategoryEntity) TableName() string {
	return "product_category"
}
