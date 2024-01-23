package customer_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type CustomerRequest struct {
	Code           string `json:"code"  validate:"required,min=3,max=20"`
	Name           string `json:"name"  validate:"required,min=3,max=200"`
	Phone          string `json:"phone" validate:"required,min=3,max=16"`
	Email          string `json:"email" validate:"required,email"`
	Address        string `json:"address" validate:"required,min=3,max=200"`
	CountryID      int64  `json:"country_id"  validate:"required"`
	ProvinceID     int64  `json:"province_id" validate:"required"`
	DistrictID     int64  `json:"district_id" validate:"required"`
	SubDistrictID  int64  `json:"sub_district_id" validate:"required"`
	UrbanVillageID int64  `json:"urban_village_id" validate:"required"`
	dto.AbstractDto
}

type ListCustomerResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerDetailResponse struct {
	ID           int64             `json:"id"`
	Code         string            `json:"code"`
	Name         string            `json:"name"`
	Phone        string            `json:"phone"`
	Email        string            `json:"email"`
	Address      string            `json:"address"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Country      dto.StructGeneral `json:"country"`
	Province     dto.StructGeneral `json:"province"`
	District     dto.StructGeneral `json:"district"`
	SubDistrict  dto.StructGeneral `json:"sub_district"`
	UrbanVillage dto.StructGeneral `json:"urban_village"`
}

func (c *CustomerRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *CustomerRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
