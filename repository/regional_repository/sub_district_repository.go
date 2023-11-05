package regional_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
)

type SubDistrictRepository interface {
	Insert(district *regional_entity.SubDistrict) model.ErrorModel
	GetByCode(code string) (regional_entity.SubDistrict, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
}
