package product_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type ProductGroupRequest struct {
	Code      string `json:"code" validate:"required,min=3,max=20"`
	Name      string `json:"name" validate:"required,min=5,max=200"`
	CompanyID int64  `json:"company_id" validate:"required"`
	Level     int64  `json:"level" validate:"required"`
	ParentID  int64  `json:"parent_id"`
	dto.AbstractDto
}

type ListProductGroupResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Level     int64     `json:"level"`
	ParentID  int64     `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DetailProductGroupResponse struct {
	ID        int64             `json:"id"`
	Code      string            `json:"code"`
	Name      string            `json:"name"`
	Level     int64             `json:"level"`
	Parent    dto.StructGeneral `json:"parent"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func (c *ProductGroupRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	mapError := common.Validation.ValidationAll(*c, contextModel)
	if c.Level > 1 && c.ParentID == 0 {
		if mapError == nil {
			mapError = make(map[string]string)
		}
		mapError["parent_id"] = "Parent is required"
	}

	return mapError
}

func (c *ProductGroupRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	if c.Level > 0 && c.ParentID == 0 {
		resultMap["parent_id"] = "Parent is required"
	}
	errMdl = c.ValidateUpdateGeneral()

	return
}
