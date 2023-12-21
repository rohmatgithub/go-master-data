package example_service

import (
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/entity"
	"go-master-data/model"
	"go-master-data/repository/example_repository"
	"go-master-data/service"
	"strings"
	"time"
)

type exampleServiceImpl struct {
	ExampleRepository example_repository.ExampleRepository
}

func NewExampleService(cpRepo example_repository.ExampleRepository) ExampleService {
	return &exampleServiceImpl{ExampleRepository: cpRepo}
}

func (cp *exampleServiceImpl) Insert(request dto.ExampleRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		return
	}

	// check by npwp
	cpDb, errMdl := cp.ExampleRepository.FetchData(entity.ExampleEntity{
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
	cpEntity := entity.ExampleEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		Code: request.Code,
		Name: request.Name,
	}

	errMdl = cp.ExampleRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *exampleServiceImpl) Update(request dto.ExampleRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		return
	}
	cpDb, errMdl := cp.ExampleRepository.FetchData(entity.ExampleEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.Example)
		return
	}

	timeNow := time.Now()
	cpEntity := entity.ExampleEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: cpDb.CreatedBy,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: cpDb.CreatedAt,
			UpdatedAt: timeNow,
			Deleted:   false,
		},
		Code: request.Code,
		Name: request.Name,
	}

	errMdl = cp.ExampleRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertError(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *exampleServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.ExampleRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []dto.ListExampleResponse
	for _, temp := range resultDB {
		data := temp.(entity.ExampleEntity)
		result = append(result, dto.ListExampleResponse{
			ID:       data.ID,
			Code:     data.Code,
			Name:     data.Name,
			Address1: data.Address1,
		})
	}
	out.Data = result
	//todo i18n
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *exampleServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.ExampleRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	out.Data = dto.DetailExampleResponse{
		ID:        dataDB.ID,
		Code:      dataDB.Code,
		Name:      dataDB.Name,
		Address1:  dataDB.Address1,
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
