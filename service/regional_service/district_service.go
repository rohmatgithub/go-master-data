package regional_service

import (
	"go-master-data/dto"
	"go-master-data/model"
)

type DistrictService interface {
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) ([]dto.DistrictListResponse, model.ErrorModel)
}
