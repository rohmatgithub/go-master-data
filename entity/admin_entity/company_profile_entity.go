package admin_entity

import "go-master-data/entity"

type CompanyProfileEntity struct {
	entity.AbstractEntity
	NPWP           string
	Name           string
	Address1       string `gorm:"column:address_1"`
	Address2       string `gorm:"column:address_2"`
	CountryID      int64
	DistrictID     int64
	SubDistrictID  int64
	UrbanVillageID int64
}

func (CompanyProfileEntity) TableName() string {
	return "company_profile"
}

type CompanyProfileDetailEntity struct {
	CompanyProfileEntity
	CountryCode      string
	CountryName      string
	DistrictCode     string
	DistrictName     string
	SubDistrictCode  string
	SubDistrictName  string
	UrbanVillageCode string
	UrbanVillageName string
}
