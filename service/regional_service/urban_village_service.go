package regional_service

import (
	"go-master-data/dto"
	"go-master-data/model"
)

type UrbanVillageService interface {
	Insert(request dto.UrbanVillageRequest) (dto.UrbanVillageResponse, model.ErrorModel)
	Import(pathStr string) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (dto.Payload, model.ErrorModel)
}
