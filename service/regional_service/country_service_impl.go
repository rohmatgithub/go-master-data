package regional_service

import (
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/dto/regional_dto"
	"go-master-data/entity/regional_entity"
	"go-master-data/model"
	"go-master-data/repository/regional_repository"
	"go-master-data/service"
)

type countryServiceImpl struct {
	CountryRepo regional_repository.CountryRepository
}

func NewCountryService(countryRepo regional_repository.CountryRepository) CountryService {
	return &countryServiceImpl{CountryRepo: countryRepo}
}

func (country *countryServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := country.CountryRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []regional_dto.CountryListResponse
	for _, temp := range resultDB {
		data := temp.(regional_entity.Country)
		result = append(result, regional_dto.CountryListResponse{
			ID:   data.ID,
			Code: data.Code,
			Name: data.Name,
		})
	}

	out.Data = result

	out.Status.Message = service.ListI18NMessage(constanta.LanguageEn)
	return
}
