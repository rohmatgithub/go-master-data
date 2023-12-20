package regional_service

import (
	"go-master-data/dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository/regional_repository"
)

type countryServiceImpl struct {
	CountryRepo regional_repository.CountryRepository
}

func NewCountryService(countryRepo regional_repository.CountryRepository) CountryService {
	return &countryServiceImpl{CountryRepo: countryRepo}
}

func (service *countryServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := service.CountryRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []dto.CountryListResponse
	for _, temp := range resultDB {
		country := temp.(regional_entity.Country)
		result = append(result, dto.CountryListResponse{
			ID:   country.ID,
			Code: country.Code,
			Name: country.Name,
		})
	}

	out.Data = result

	// todo i18n
	out.Status.Message = "Berhasil ambil data list"
	return
}
