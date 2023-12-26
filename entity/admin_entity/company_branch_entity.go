package admin_entity

import "go-master-data/entity"

type CompanyBranchEntity struct {
	entity.AbstractEntity
	Code             string
	CompanyID        int64
	CompanyProfileID int64
}

type CompanyBranchDetailEntity struct {
	CompanyBranchEntity
	CompanyProfile CompanyProfileEntity
	CompanyCode    string
	CompanyName    string
}

func (CompanyBranchEntity) TableName() string {
	return "company_branch"
}
