package regional_service

import (
	"go-master-data/dto"
	"go-master-data/model"
)

type CountryService interface {
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (dto.Payload, model.ErrorModel)
}
