package dto

import "go-master-data/common"

type CompanyProfileRequest struct {
	ID       int    `json:"id"`
	NPWP     string `json:"npwp" validate:"required,numeric,min=16,max=16"`
	Name     string `json:"name" validate:"required,min=3,max=200"`
	Address1 string `json:"address_1" validate:"required,min=3,max=250"`
	Address2 string `json:"address_2" validate:"max=250"`
}

func (c *CompanyProfileRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}
