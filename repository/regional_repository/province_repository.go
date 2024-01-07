package regional_repository

import (
	"go-master-data/dto"
	"go-master-data/model"
)

type ProvinceRepository interface {
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
}
