package product_entity

import "go-master-data/entity"

type ProductGroupEntity struct {
	entity.AbstractEntity
	Code      string
	Name      string
	CompanyID int64
	Level     int64
	ParentID  int64
}

type ProductGroupDetailEntity struct {
	ProductGroupEntity
	ParentCode string
	ParentName string
}

func (ProductGroupEntity) TableName() string {
	return "product_group_hierarchy"
}
