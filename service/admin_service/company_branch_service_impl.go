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

type companyBranchServiceImpl struct {
	CompanyBranchRepository admin_repository.CompanyBranchRepository
}

func NewCompanyBranchService(cpRepo admin_repository.CompanyBranchRepository) CompanyBranchService {
	return &companyBranchServiceImpl{CompanyBranchRepository: cpRepo}
}

func (cp *companyBranchServiceImpl) Insert(request admin_dto.CompanyBranchRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// check by npwp
	cpDb, errMdl := cp.CompanyBranchRepository.FetchData(admin_entity.CompanyBranchEntity{
		Code: request.Code,
	})
	if errMdl.Error != nil {
		return
	}

	if cpDb.ID > 0 {
		errMdl = model.GenerateHasUsedDataError(constanta.Code)
		return
	}
	// insert
	timeNow := time.Now()
	cpEntity := admin_entity.CompanyBranchEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID:        request.CompanyID,
		CompanyProfileID: request.CompanyProfileID,
		Code:             request.Code,
	}

	errMdl = cp.CompanyBranchRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyBranchServiceImpl) Update(request admin_dto.CompanyBranchRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}
	cpDb, errMdl := cp.CompanyBranchRepository.FetchData(admin_entity.CompanyBranchEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.CompanyBranch)
		return
	}

	timeNow := time.Now()
	cpEntity := admin_entity.CompanyBranchEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: cpDb.CreatedBy,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: cpDb.CreatedAt,
			UpdatedAt: timeNow,
			Deleted:   false,
		},
		CompanyID:        request.CompanyID,
		CompanyProfileID: request.CompanyProfileID,
		Code:             request.Code,
	}

	errMdl = cp.CompanyBranchRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorCompanyBranch(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *companyBranchServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyBranchRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []admin_dto.ListCompanyBranchResponse
	for _, temp := range resultDB {
		data := temp.(admin_entity.CompanyBranchDetailEntity)
		result = append(result, admin_dto.ListCompanyBranchResponse{
			ID:        data.ID,
			Code:      data.Code,
			Name:      data.CompanyProfile.Name,
			Address1:  data.CompanyProfile.Address1,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	out.Data = result
	//todo i18n
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyBranchServiceImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyBranchRepository.Count(searchParam)
	if errMdl.Error != nil {
		return
	}

	out.Data = resultDB
	out.Status.Message = service.CountI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyBranchServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.CompanyBranchRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = admin_dto.DetailCompanyBranchResponse{
		ID:               dataDB.ID,
		CompanyProfileID: dataDB.CompanyProfileID,
		CompanyID:        dataDB.CompanyID,
		CompanyCode:      dataDB.CompanyCode,
		CompanyName:      dataDB.CompanyName,
		NPWP:             dataDB.CompanyProfile.NPWP,
		Code:             dataDB.Code,
		Name:             dataDB.CompanyProfile.Name,
		Address1:         dataDB.CompanyProfile.Address1,
		CreatedAt:        dataDB.CreatedAt,
		UpdatedAt:        dataDB.UpdatedAt,
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorCompanyBranch(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "company_branch_code_key") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
