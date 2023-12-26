package product_service

import (
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/dto/product_dto"
	"go-master-data/entity"
	"go-master-data/entity/product_entity"
	"go-master-data/model"
	"go-master-data/repository/product_repository"
	"go-master-data/service"
	"strings"
	"time"
)

type productGroupServiceImpl struct {
	ProductGroupRepository product_repository.ProductGroupRepository
}

func NewProductGroupService(cpRepo product_repository.ProductGroupRepository) ProductGroupService {
	return &productGroupServiceImpl{ProductGroupRepository: cpRepo}
}

func (cp *productGroupServiceImpl) Insert(request product_dto.ProductGroupRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		return
	}

	// check by npwp
	cpDb, errMdl := cp.ProductGroupRepository.FetchData(product_entity.ProductGroupEntity{
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
	cpEntity := product_entity.ProductGroupEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID: request.CompanyID,
		Code:      request.Code,
		Name:      request.Name,
		Level:     request.Level,
		ParentID:  request.ParentID,
	}

	errMdl = cp.ProductGroupRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productGroupServiceImpl) Update(request product_dto.ProductGroupRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		return
	}
	cpDb, errMdl := cp.ProductGroupRepository.FetchData(product_entity.ProductGroupEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ProductGroup)
		return
	}

	timeNow := time.Now()
	cpEntity := product_entity.ProductGroupEntity{
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
		Level:     request.Level,
		ParentID:  request.ParentID,
	}

	errMdl = cp.ProductGroupRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorProductGroup(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *productGroupServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	resultDB, errMdl := cp.ProductGroupRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []product_dto.ListProductGroupResponse
	for _, temp := range resultDB {
		data := temp.(product_entity.ProductGroupEntity)
		result = append(result, product_dto.ListProductGroupResponse{
			ID:       data.ID,
			Code:     data.Code,
			Name:     data.Name,
			Level:    data.Level,
			ParentID: data.ParentID,
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productGroupServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.ProductGroupRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = product_dto.DetailProductGroupResponse{
		ID:        dataDB.ID,
		Code:      dataDB.Code,
		Name:      dataDB.Name,
		Level:     dataDB.Level,
		ParentID:  dataDB.ParentID,
		CreatedAt: dataDB.CreatedAt,
		UpdatedAt: dataDB.UpdatedAt,
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorProductGroup(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "uq_productcategory_code_company") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
