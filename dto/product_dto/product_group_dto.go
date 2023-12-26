package product_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type ProductGroupRequest struct {
	Code       string `json:"code" validate:"required,min=3,max=20"`
	Name       string `json:"name" validate:"required,min=5,max=200"`
	CompanyID  int64  `json:"company_id" validate:"required"`
	DivisionID int64  `json:"division_id" validate:"required"`
	Level      int64  `json:"level" validate:"required"`
	ParentID   int64  `json:"parent_id"`
	dto.AbstractDto
}

type ListProductGroupResponse struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Level    int64  `json:"level"`
	ParentID int64  `json:"parent_id"`
}

type DetailProductGroupResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Level     int64     `json:"level"`
	ParentID  int64     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *ProductGroupRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *ProductGroupRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
