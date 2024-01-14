package admin_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type CompanyRequest struct {
	Code             string `json:"code" validate:"required,min=3,max=20"`
	CompanyProfileID int64  `json:"company_profile_id" validate:"required"`
	dto.AbstractDto
}

type ListCompanyResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address1  string    `json:"address_1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailCompanyResponse struct {
	ID               int64     `json:"id"`
	CompanyProfileID int64     `json:"company_profile_id"`
	Code             string    `json:"code"`
	Name             string    `json:"name"`
	NPWP             string    `json:"npwp"`
	Address1         string    `json:"address_1"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (c *CompanyRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *CompanyRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
