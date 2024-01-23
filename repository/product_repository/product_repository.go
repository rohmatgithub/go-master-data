package product_repository

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/entity/product_entity"
	"go-master-data/model"
)

type ProductRepository interface {
	Insert(entity *product_entity.ProductEntity) model.ErrorModel
	Update(entity *product_entity.ProductEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result []interface{}, errMdl model.ErrorModel)
	View(id int64) (product_entity.ProductDetailEntity, model.ErrorModel)
	FetchData(entity product_entity.ProductEntity) (product_entity.ProductEntity, model.ErrorModel)
	Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result int64, errMdl model.ErrorModel)
}
