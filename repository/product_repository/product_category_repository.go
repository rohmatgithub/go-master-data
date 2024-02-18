package product_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/product_entity"
	"go-master-data/model"
)

type ProductCategoryRepository interface {
	Insert(entity *product_entity.ProductCategoryEntity) model.ErrorModel
	Update(entity *product_entity.ProductCategoryEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
	View(id int64) (product_entity.ProductCategoryEntity, model.ErrorModel)
	FetchData(entity product_entity.ProductCategoryEntity) (product_entity.ProductCategoryEntity, model.ErrorModel)
	Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel)
}
