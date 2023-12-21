package example_service

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
)

type ExampleService interface {
	Insert(request dto.ExampleRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Update(request dto.ExampleRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	ViewDetail(id int64, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
}
