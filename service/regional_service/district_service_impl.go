package regional_service

import (
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository/regional_repository"
)

type districtServiceImpl struct {
	DistrictRepo regional_repository.DistrictRepository
}

func NewDistrictService(districtRepo regional_repository.DistrictRepository) DistrictService {
	return &districtServiceImpl{DistrictRepo: districtRepo}
}

func (service *districtServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := service.DistrictRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []dto.DistrictListResponse
	for _, temp := range resultDB {
		district := temp.(regional_entity.District)
		result = append(result, dto.DistrictListResponse{
			ID:       district.ID,
			ParentID: district.ParentID,
			Code:     district.Code,
			Name:     district.Name,
		})
	}

	out.Data = result

	// todo i18n
	out.Status.Message = "Berhasil ambil data list"
	return
}
