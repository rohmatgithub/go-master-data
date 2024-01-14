package admin_entity

import "go-master-data/entity"

type CompanyDivisionEntity struct {
	entity.AbstractEntity
	Code      string
	Name      string
	CompanyID int64
}

type CompanyDivisionDetailEntity struct {
	CompanyDivisionEntity
	CompanyCode string
	CompanyName string
}

func (CompanyDivisionEntity) TableName() string {
	return "company_division"
}
