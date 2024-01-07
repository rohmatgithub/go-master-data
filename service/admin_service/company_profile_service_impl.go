package admin_service

import (
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/dto/admin_dto"
	"go-master-data/entity"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
	"go-master-data/repository/admin_repository"
	"go-master-data/service"
	"strings"
	"time"
)

type companyProfileServiceImpl struct {
	CompanyProfileRepository admin_repository.CompanyProfileRepository
}

func NewCompanyProfileService(cpRepo admin_repository.CompanyProfileRepository) CompanyProfileService {
	return &companyProfileServiceImpl{CompanyProfileRepository: cpRepo}
}

func (cp *companyProfileServiceImpl) Insert(request admin_dto.CompanyProfileRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// check by npwp
	cpDb, errMdl := cp.CompanyProfileRepository.FetchData(admin_entity.CompanyProfileEntity{
		NPWP: request.NPWP,
	})
	if errMdl.Error != nil {
		return
	}

	if cpDb.ID > 0 {
		errMdl = model.GenerateHasUsedDataError(constanta.Npwp)
		return
	}
	// insert
	timeNow := time.Now()
	cpEntity := admin_entity.CompanyProfileEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		NPWP:           request.NPWP,
		Name:           request.Name,
		Address1:       request.Address1,
		Address2:       request.Address2,
		CountryID:      request.CountryID,
		ProvinceID:     request.ProvinceID,
		DistrictID:     request.DistrictID,
		SubDistrictID:  request.SubDistrictID,
		UrbanVillageID: request.UrbanVillageID,
	}

	errMdl = cp.CompanyProfileRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyProfileServiceImpl) Update(request admin_dto.CompanyProfileRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}
	cpDb, errMdl := cp.CompanyProfileRepository.FetchData(admin_entity.CompanyProfileEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.CompanyProfile)
		return
	}

	timeNow := time.Now()
	cpEntity := admin_entity.CompanyProfileEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: cpDb.CreatedBy,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: cpDb.CreatedAt,
			UpdatedAt: timeNow,
			Deleted:   false,
		},
		NPWP:           request.NPWP,
		Name:           request.Name,
		Address1:       request.Address1,
		Address2:       request.Address2,
		CountryID:      request.CountryID,
		ProvinceID:     request.ProvinceID,
		DistrictID:     request.DistrictID,
		SubDistrictID:  request.SubDistrictID,
		UrbanVillageID: request.UrbanVillageID,
	}

	errMdl = cp.CompanyProfileRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertError(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *companyProfileServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyProfileRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []admin_dto.ListCompanyProfileResponse
	for _, temp := range resultDB {
		data := temp.(admin_entity.CompanyProfileEntity)
		result = append(result, admin_dto.ListCompanyProfileResponse{
			ID:        data.ID,
			NPWP:      data.NPWP,
			Name:      data.Name,
			Address1:  data.Address1,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyProfileServiceImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyProfileRepository.Count(searchParam)
	if errMdl.Error != nil {
		return
	}

	out.Data = resultDB
	out.Status.Message = service.CountI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyProfileServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.CompanyProfileRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = admin_dto.DetailCompanyProfile{
		ID:       dataDB.ID,
		NPWP:     dataDB.NPWP,
		Name:     dataDB.Name,
		Address1: dataDB.Address1,
		Address2: dataDB.Address2,
		Country: dto.StructGeneral{
			ID:   dataDB.CountryID,
			Code: dataDB.CountryCode,
			Name: dataDB.CountryName,
		},
		District: dto.StructGeneral{
			ID:   dataDB.DistrictID,
			Code: dataDB.DistrictCode,
			Name: dataDB.DistrictName,
		},
		SubDistrict: dto.StructGeneral{
			ID:   dataDB.SubDistrictID,
			Code: dataDB.SubDistrictCode,
			Name: dataDB.SubDistrictName,
		},
		UrbanVillage: dto.StructGeneral{
			ID:   dataDB.UrbanVillageID,
			Code: dataDB.UrbanVillageCode,
			Name: dataDB.UrbanVillageName,
		},
		CreatedAt: dataDB.CreatedAt,
		UpdatedAt: dataDB.UpdatedAt,
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertError(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "company_profile_npwp_key") {
		return model.GenerateHasUsedDataError(constanta.Npwp)
	}
	return
}
