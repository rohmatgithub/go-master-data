package admin_entity

import "go-master-data/entity"

type CompanyEntity struct {
	entity.AbstractEntity
	Code             string
	CompanyProfileID int64
}

type CompanyDetailEntity struct {
	CompanyEntity
	CompanyProfile CompanyProfileEntity
}

func (CompanyEntity) TableName() string {
	return "company"
}
