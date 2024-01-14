package admin_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type CompanyProfileRequest struct {
	NPWP           string `json:"npwp" validate:"required,numeric,min=16,max=16"`
	Name           string `json:"name" validate:"required,min=3,max=200"`
	Address1       string `json:"address_1" validate:"required,min=3,max=250"`
	Address2       string `json:"address_2" validate:"max=250"`
	CountryID      int64  `json:"country_id" validate:"required"`
	ProvinceID     int64  `json:"province_id" validate:"required"`
	DistrictID     int64  `json:"district_id" validate:"required"`
	SubDistrictID  int64  `json:"sub_district_id" validate:"required"`
	UrbanVillageID int64  `json:"urban_village_id" validate:"required"`
	dto.AbstractDto
}

type ListCompanyProfileResponse struct {
	ID        int64     `json:"id"`
	NPWP      string    `json:"npwp"`
	Name      string    `json:"name"`
	Address1  string    `json:"address_1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailCompanyProfile struct {
	ID           int64             `json:"id"`
	NPWP         string            `json:"npwp"`
	Name         string            `json:"name"`
	Address1     string            `json:"address_1"`
	Address2     string            `json:"address_2"`
	Country      dto.StructGeneral `json:"country"`
	Province     dto.StructGeneral `json:"province"`
	District     dto.StructGeneral `json:"district"`
	SubDistrict  dto.StructGeneral `json:"sub_district"`
	UrbanVillage dto.StructGeneral `json:"urban_village"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
}

func (c *CompanyProfileRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *CompanyProfileRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
