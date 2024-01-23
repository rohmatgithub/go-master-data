package product_dto

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type ProductRequest struct {
	DivisionID   int64   `json:"division_id" validate:"required"`
	CategoryID   int64   `json:"category_id" validate:"required"`
	GroupID      int64   `json:"group_id" validate:"required"`
	Code         string  `json:"code" validate:"required,min=3,max=20"`
	Name         string  `json:"name" validate:"required,min=3,max=150"`
	SellingPrice float64 `json:"selling_price" validate:"required"`
	BuyingPrice  float64 `json:"buying_price" validate:"required"`
	Uom1         string  `json:"uom_1" validate:"required,min=1,max=10"`
	Uom2         string  `json:"uom_2" validate:"required,min=1,max=10"`
	Conv1To2     int32   `json:"conv1_to2" validate:"required"`
	dto.AbstractDto
}

type ListProductResponse struct {
	ID        int64             `json:"id"`
	Code      string            `json:"code"`
	Name      string            `json:"name"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	Division  dto.StructGeneral `json:"division"`
}

type DetailProductResponse struct {
	ID           int64             `json:"id"`
	Code         string            `json:"code"`
	Name         string            `json:"name"`
	SellingPrice float64           `json:"sell_price"`
	BuyingPrice  float64           `json:"buy_price"`
	Uom1         string            `json:"uom1"`
	Uom2         string            `json:"uom2"`
	Conv1To2     int32             `json:"conv_1_to_2"`
	Division     dto.StructGeneral `json:"division"`
	Category     dto.StructGeneral `json:"category"`
	Group        dto.StructGeneral `json:"group"`
	UpdatedAt    time.Time         `json:"updated_at"`
	CreatedAt    time.Time         `json:"created_at"`
}

func (c *ProductRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *ProductRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.ValidateUpdateGeneral()

	return
}
