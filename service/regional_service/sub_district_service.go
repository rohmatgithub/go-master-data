package regional_service

import (
	"go-master-data/dto"
	"go-master-data/model"
)

type SubDistrictService interface {
	Insert(request dto.SubDistrictRequest) (dto.SubDistrictResponse, model.ErrorModel)
	Import(pathStr string) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (dto.Payload, model.ErrorModel)
}
