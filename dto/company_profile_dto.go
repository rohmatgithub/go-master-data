package dto

import (
	"go-master-data/common"
	"go-master-data/model"
)

type CompanyProfileRequest struct {
	NPWP           string `json:"npwp" validate:"required,numeric,min=16,max=16"`
	Name           string `json:"name" validate:"required,min=3,max=200"`
	Address1       string `json:"address_1" validate:"required,min=3,max=250"`
	Address2       string `json:"address_2" validate:"max=250"`
	CountryID      int64  `json:"country_id" validate:"required"`
	DistrictID     int64  `json:"district_id" validate:"required"`
	SubDistrictID  int64  `json:"sub_district_id" validate:"required"`
	UrbanVillageID int64  `json:"urban_village_id" validate:"required"`
	AbstractDto
}

func (c *CompanyProfileRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *CompanyProfileRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.validateUpdate()

	return
}
