package regional_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
)

type UrbanVillageRepository interface {
	Insert(urbanVillage *regional_entity.UrbanVillage) model.ErrorModel
	GetByCode(code string) (regional_entity.UrbanVillage, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
}
