package admin_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type CompanyBranchRequest struct {
	Code             string `json:"code" validate:"required,min=3,max=20"`
	CompanyID        int64  `json:"company_id" validate:"required"`
	CompanyProfileID int64  `json:"company_profile_id" validate:"required"`
	dto.AbstractDto
}

type ListCompanyBranchResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address1  string    `json:"address_1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailCompanyBranchResponse struct {
	ID               int64     `json:"id"`
	CompanyProfileID int64     `json:"company_profile_id"`
	CompanyID        int64     `json:"company_id"`
	CompanyCode      string    `json:"company_code"`
	CompanyName      string    `json:"company_name"`
	Code             string    `json:"code"`
	NPWP             string    `json:"npwp"`
	Name             string    `json:"name"`
	Address1         string    `json:"address_1"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (c *CompanyBranchRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *CompanyBranchRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
