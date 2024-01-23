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
	"go-master-data/service/admin_service"
	"strings"
	"time"
)

type companyServiceImpl struct {
	CompanyRepository admin_repository.CompanyRepository
}

func NewCompanyService(cpRepo admin_repository.CompanyRepository) admin_service.CompanyService {
	return &companyServiceImpl{CompanyRepository: cpRepo}
}

func (cp *companyServiceImpl) Insert(request admin_dto.CompanyRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// check by npwp
	cpDb, errMdl := cp.CompanyRepository.FetchData(admin_entity.CompanyEntity{
		Code: request.Code,
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
	cpEntity := admin_entity.CompanyEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyProfileID: request.CompanyProfileID,
		Code:             request.Code,
	}

	errMdl = cp.CompanyRepository.Insert(&cpEntity)

	//out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyServiceImpl) Update(request admin_dto.CompanyRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}
	cpDb, errMdl := cp.CompanyRepository.FetchData(admin_entity.CompanyEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.Company)
		return
	}

	timeNow := time.Now()
	cpEntity := admin_entity.CompanyEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: cpDb.CreatedBy,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: cpDb.CreatedAt,
			UpdatedAt: timeNow,
			Deleted:   false,
		},
		CompanyProfileID: request.CompanyProfileID,
		Code:             request.Code,
	}

	errMdl = cp.CompanyRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorCompany(errMdl)
	}
	//out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *companyServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []admin_dto.ListCompanyResponse
	for _, temp := range resultDB {
		data := temp.(admin_entity.CompanyDetailEntity)
		result = append(result, admin_dto.ListCompanyResponse{
			ID:        data.ID,
			Code:      data.Code,
			Name:      data.CompanyProfile.Name,
			Address1:  data.CompanyProfile.Address1,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyServiceImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyRepository.Count(searchParam)
	if errMdl.Error != nil {
		return
	}

	out.Data = resultDB
	//out.Status.Message = service.CountI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.CompanyRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = admin_dto.DetailCompanyResponse{
		ID:               dataDB.ID,
		CompanyProfileID: dataDB.CompanyProfileID,
		Code:             dataDB.Code,
		NPWP:             dataDB.CompanyProfile.NPWP,
		Name:             dataDB.CompanyProfile.Name,
		Address1:         dataDB.CompanyProfile.Address1,
		CreatedAt:        dataDB.CreatedAt,
		UpdatedAt:        dataDB.UpdatedAt,
	}

	//out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorCompany(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "company_code_key") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
