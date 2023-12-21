package dto

import (
	"go-master-data/common"
	"go-master-data/model"
	"time"
)

type ExampleRequest struct {
	Code      string `json:"code" validate:"required,min=16,max=16"`
	Name      string `json:"name" validate:"required,min=16,max=16"`
	ForeignID int64  `json:"foreign_id" validate:"required"`
	AbstractDto
}

type ListExampleResponse struct {
	ID       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Address1 string `json:"address_1"`
}

type DetailExampleResponse struct {
	ID        int64     `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Address1  string    `json:"address_1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *ExampleRequest) ValidateInsert(contextModel *common.ContextModel) map[string]string {
	return common.Validation.ValidationAll(*c, contextModel)
}

func (c *ExampleRequest) ValidateUpdate(contextModel *common.ContextModel) (resultMap map[string]string, errMdl model.ErrorModel) {
	resultMap = common.Validation.ValidationAll(*c, contextModel)

	errMdl = c.validateUpdate()

	return
}
