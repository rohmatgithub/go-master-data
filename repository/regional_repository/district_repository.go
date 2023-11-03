package regional_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
)

type DistrictRepository interface {
	Insert(district *regional_entity.District) model.ErrorModel
	GetByCode(code string) (regional_entity.District, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
}
