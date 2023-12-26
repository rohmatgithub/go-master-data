package regional_service

import (
	"go-master-data/dto"
	"go-master-data/dto/regional_dto"
	"go-master-data/model"
)

type UrbanVillageService interface {
	Insert(request regional_dto.UrbanVillageRequest) (regional_dto.UrbanVillageResponse, model.ErrorModel)
	Import(pathStr string) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (dto.Payload, model.ErrorModel)
}
