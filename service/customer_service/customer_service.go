package customer_service

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/dto/customer_dto"
	"go-master-data/model"
)

type CustomerService interface {
	Insert(request customer_dto.CustomerRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Update(request customer_dto.CustomerRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	ViewDetail(id int64, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
}
