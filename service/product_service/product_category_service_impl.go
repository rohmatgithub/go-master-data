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
	"strconv"
	"strings"
	"time"
)

type productCategoryServiceImpl struct {
	ProductCategoryRepository product_repository.ProductCategoryRepository
}

func NewProductCategoryService(cpRepo product_repository.ProductCategoryRepository) ProductCategoryService {
	return &productCategoryServiceImpl{ProductCategoryRepository: cpRepo}
}

func (cp *productCategoryServiceImpl) Insert(request product_dto.ProductCategoryRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	request.CompanyID = ctxModel.AuthAccessTokenModel.CompanyID
	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// check by code
	cpDb, errMdl := cp.ProductCategoryRepository.FetchData(product_entity.ProductCategoryEntity{
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
	cpEntity := product_entity.ProductCategoryEntity{
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

	errMdl = cp.ProductCategoryRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productCategoryServiceImpl) Update(request product_dto.ProductCategoryRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	request.CompanyID = ctxModel.AuthAccessTokenModel.CompanyID
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		return
	}
	cpDb, errMdl := cp.ProductCategoryRepository.FetchData(product_entity.ProductCategoryEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ProductCategory)
		return
	}

	timeNow := time.Now()
	cpEntity := product_entity.ProductCategoryEntity{
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

	errMdl = cp.ProductCategoryRepository.Update(&cpEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorProductCategory(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *productCategoryServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	})

	resultDB, errMdl := cp.ProductCategoryRepository.List(dtoList, searchParam)
	if errMdl.Error != nil {
		return
	}

	var result []product_dto.ListProductCategoryResponse
	for _, temp := range resultDB {
		data := temp.(product_entity.ProductCategoryEntity)
		result = append(result, product_dto.ListProductCategoryResponse{
			ID:        data.ID,
			Code:      data.Code,
			Name:      data.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productCategoryServiceImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	searchParam = append(searchParam, dto.SearchByParam{
		SearchKey:      "company_id",
		SearchOperator: "eq",
		SearchValue:    strconv.Itoa(int(ctxModel.AuthAccessTokenModel.CompanyID)),
	})
	resultDB, errMdl := cp.ProductCategoryRepository.Count(searchParam)
	if errMdl.Error != nil {
		return
	}
	out.Data = resultDB
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productCategoryServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.ProductCategoryRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = product_dto.DetailProductCategoryResponse{
		ID:        dataDB.ID,
		Code:      dataDB.Code,
		Name:      dataDB.Name,
		CreatedAt: dataDB.CreatedAt,
		UpdatedAt: dataDB.UpdatedAt,
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorProductCategory(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "uq_productcategory_code_company") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
