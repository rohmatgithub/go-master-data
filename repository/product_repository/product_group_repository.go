package product_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/product_entity"
	"go-master-data/model"
)

type ProductGroupRepository interface {
	Insert(entity *product_entity.ProductGroupEntity) model.ErrorModel
	Update(entity *product_entity.ProductGroupEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
	View(id int64) (product_entity.ProductGroupDetailEntity, model.ErrorModel)
	FetchData(entity product_entity.ProductGroupEntity) (product_entity.ProductGroupEntity, model.ErrorModel)
}
