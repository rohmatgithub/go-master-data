package regional_service

import (
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/dto/regional_dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository/regional_repository"
	"go-master-data/service"
	"strconv"
)

type districtServiceImpl struct {
	DistrictRepo regional_repository.DistrictRepository
}

func NewDistrictService(districtRepo regional_repository.DistrictRepository) DistrictService {
	return &districtServiceImpl{DistrictRepo: districtRepo}
}

func (district *districtServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {

	parentID := 0
	for _, param := range searchParam {
		if param.SearchKey == "parent_id" {
			parentID, _ = strconv.Atoi(param.SearchValue)
			break
		}
	}
	if parentID == 0 {
		errMdl = model.GenerateEmptyFieldError(constanta.ParentID)
		return
	}
	resultDB, errMdl := district.DistrictRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []regional_dto.DistrictListResponse
	for _, temp := range resultDB {
		data := temp.(regional_entity.District)
		result = append(result, regional_dto.DistrictListResponse{
			ID:       data.ID,
			ParentID: data.ParentID,
			Code:     data.Code,
			Name:     data.Name,
		})
	}

	out.Data = result

	out.Status.Message = service.ListI18NMessage(constanta.LanguageEn)
	return
}
