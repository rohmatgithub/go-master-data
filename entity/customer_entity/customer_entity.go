package customer_entity

import (
	"database/sql"
	"go-master-data/entity"
)

type CustomerEntity struct {
	CompanyID      sql.NullInt64
	BranchID       sql.NullInt64
	Code           sql.NullString
	Name           sql.NullString
	Phone          sql.NullString
	Email          sql.NullString
	Address        sql.NullString
	CountryID      sql.NullInt64
	ProvinceID     sql.NullInt64
	DistrictID     sql.NullInt64
	SubDistrictID  sql.NullInt64
	UrbanVillageID sql.NullInt64
	entity.AbstractEntity
}

type CustomerDetailEntity struct {
	CustomerEntity
	CountryCode      sql.NullString
	CountryName      sql.NullString
	ProvinceCode     sql.NullString
	ProvinceName     sql.NullString
	DistrictCode     sql.NullString
	DistrictName     sql.NullString
	SubDistrictCode  sql.NullString
	SubDistrictName  sql.NullString
	UrbanVillageCode sql.NullString
	UrbanVillageName sql.NullString
}

func (CustomerEntity) TableName() string {
	return "customer"
}
