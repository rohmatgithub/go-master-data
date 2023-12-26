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

type companyDivisionServiceImpl struct {
	CompanyDivisionRepository admin_repository.CompanyDivisionRepository
}

func NewCompanyDivisionService(cpRepo admin_repository.CompanyDivisionRepository) CompanyDivisionService {
	return &companyDivisionServiceImpl{CompanyDivisionRepository: cpRepo}
}

func (cp *companyDivisionServiceImpl) Insert(request admin_dto.CompanyDivisionRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		return
	}

	// check by npwp
	cpDb, errMdl := cp.CompanyDivisionRepository.FetchData(admin_entity.CompanyDivisionEntity{
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
	cpEntity := admin_entity.CompanyDivisionEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID: request.CompanyID,
		Code:      request.Code,
		Name:      request.Name,
	}

	errMdl = cp.CompanyDivisionRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyDivisionServiceImpl) Update(request admin_dto.CompanyDivisionRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		return
	}
	cpDb, errMdl := cp.CompanyDivisionRepository.FetchData(admin_entity.CompanyDivisionEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.CompanyDivision)
		return
	}

	timeNow := time.Now()
	cpEntity := admin_entity.CompanyDivisionEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: cpDb.CreatedBy,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: cpDb.CreatedAt,
			UpdatedAt: timeNow,
			Deleted:   false,
		},
		CompanyID: request.CompanyID,
		Code:      request.Code,
		Name:      request.Name,
	}

	errMdl = cp.CompanyDivisionRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorCompanyDivision(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *companyDivisionServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.CompanyDivisionRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []admin_dto.ListCompanyDivisionResponse
	for _, temp := range resultDB {
		data := temp.(admin_entity.CompanyDivisionEntity)
		result = append(result, admin_dto.ListCompanyDivisionResponse{
			ID:   data.ID,
			Code: data.Code,
			Name: data.Name,
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyDivisionServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.CompanyDivisionRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = admin_dto.DetailCompanyDivisionResponse{
		ID:        dataDB.ID,
		Code:      dataDB.Code,
		Name:      dataDB.Name,
		CreatedAt: dataDB.CreatedAt,
		UpdatedAt: dataDB.UpdatedAt,
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorCompanyDivision(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "uq_companydivision_code_company") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
