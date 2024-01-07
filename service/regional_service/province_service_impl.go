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

type provinceServiceImpl struct {
	ProvinceRepo regional_repository.ProvinceRepository
}

func NewProvinceService(provinceRepo regional_repository.ProvinceRepository) ProvinceService {
	return &provinceServiceImpl{ProvinceRepo: provinceRepo}
}

func (province *provinceServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {

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
	resultDB, errMdl := province.ProvinceRepo.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []regional_dto.ProvinceListResponse
	for _, temp := range resultDB {
		data := temp.(regional_entity.Province)
		result = append(result, regional_dto.ProvinceListResponse{
			ID:   data.ID,
			Code: data.Code,
			Name: data.Name,
		})
	}

	out.Data = result

	out.Status.Message = service.ListI18NMessage(constanta.LanguageEn)
	return
}
