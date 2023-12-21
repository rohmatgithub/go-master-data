package example_repository

import (
	"go-master-data/dto"
	"go-master-data/entity"
	"go-master-data/model"
)

type ExampleRepository interface {
	Insert(entity *entity.ExampleEntity) model.ErrorModel
	Update(entity *entity.ExampleEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
	View(id int64) (entity.ExampleDetailEntity, model.ErrorModel)
	FetchData(entity entity.ExampleEntity) (entity.ExampleEntity, model.ErrorModel)
}
