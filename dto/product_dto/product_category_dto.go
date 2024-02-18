package product_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type ProductCategoryRequest struct {
	Code      string `json:"code" validate:"required,min=3,max=20"`
	Name      string `json:"name" validate:"required,min=5,max=200"`
	CompanyID int64  `json:"company_id" validate:"required"`
	dto.AbstractDto
}

type ListProductCategoryResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailProductCategoryResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *ProductCategoryRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *ProductCategoryRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
