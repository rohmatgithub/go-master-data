package product_service

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/dto/product_dto"
	"go-master-data/model"
)

type ProductService interface {
	Insert(request product_dto.ProductRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Update(request product_dto.ProductRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	ViewDetail(id int64, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
}
