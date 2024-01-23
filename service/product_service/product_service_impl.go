package product_service

import (
	"database/sql"
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

type productServiceImpl struct {
	ProductRepository product_repository.ProductRepository
}

func NewProductService(cpRepo product_repository.ProductRepository) ProductService {
	return &productServiceImpl{ProductRepository: cpRepo}
}

func (cp *productServiceImpl) Insert(request product_dto.ProductRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// check by code
	cpDb, errMdl := cp.ProductRepository.FetchData(product_entity.ProductEntity{
		Code: sql.NullString{String: request.Code, Valid: true},
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
	pEntity := product_entity.ProductEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID:    sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.CompanyID, Valid: true},
		DivisionID:   sql.NullInt64{Int64: request.DivisionID, Valid: true},
		CategoryID:   sql.NullInt64{Int64: request.CategoryID, Valid: true},
		GroupID:      sql.NullInt64{Int64: request.GroupID, Valid: true},
		Code:         sql.NullString{String: request.Code, Valid: true},
		Name:         sql.NullString{String: request.Name, Valid: true},
		SellingPrice: sql.NullFloat64{Float64: request.SellingPrice, Valid: true},
		BuyingPrice:  sql.NullFloat64{Float64: request.BuyingPrice, Valid: true},
		Uom1:         sql.NullString{String: request.Uom1, Valid: true},
		Uom2:         sql.NullString{String: request.Uom2, Valid: true},
		Conv1To2:     sql.NullInt32{Int32: request.Conv1To2, Valid: true},
	}

	errMdl = cp.ProductRepository.Insert(&pEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productServiceImpl) Update(request product_dto.ProductRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		return
	}
	cpDb, errMdl := cp.ProductRepository.FetchData(product_entity.ProductEntity{
		AbstractEntity: entity.AbstractEntity{ID: request.ID},
	})
	if errMdl.Error != nil {
		return
	}
	if cpDb.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.Product)
		return
	}

	timeNow := time.Now()
	pEntity := product_entity.ProductEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID:    sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.CompanyID, Valid: true},
		DivisionID:   sql.NullInt64{Int64: request.DivisionID, Valid: true},
		CategoryID:   sql.NullInt64{Int64: request.CategoryID, Valid: true},
		GroupID:      sql.NullInt64{Int64: request.GroupID, Valid: true},
		Code:         sql.NullString{String: request.Code, Valid: true},
		Name:         sql.NullString{String: request.Name, Valid: true},
		SellingPrice: sql.NullFloat64{Float64: request.SellingPrice, Valid: true},
		BuyingPrice:  sql.NullFloat64{Float64: request.BuyingPrice, Valid: true},
		Uom1:         sql.NullString{String: request.Uom1, Valid: true},
		Uom2:         sql.NullString{String: request.Uom2, Valid: true},
		Conv1To2:     sql.NullInt32{Int32: request.Conv1To2, Valid: true},
	}

	errMdl = cp.ProductRepository.Update(&pEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorProduct(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *productServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := cp.ProductRepository.List(dtoList, searchParam, ctxModel)
	if errMdl.Error != nil {
		return
	}

	var result []product_dto.ListProductResponse
	for _, temp := range resultDB {
		data := temp.(product_entity.ProductDetailEntity)
		result = append(result, product_dto.ListProductResponse{
			ID:        data.ID,
			Code:      data.Code.String,
			Name:      data.Name.String,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
			Division: dto.StructGeneral{
				ID:   data.DivisionID.Int64,
				Code: data.DivisionCode.String,
				Name: data.DivisionName.String,
			},
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productServiceImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := cp.ProductRepository.Count(searchParam, ctxModel)
	if errMdl.Error != nil {
		return
	}
	out.Data = resultDB
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *productServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.ProductRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = product_dto.DetailProductResponse{
		ID:           dataDB.ID,
		Code:         dataDB.Code.String,
		Name:         dataDB.Name.String,
		SellingPrice: dataDB.SellingPrice.Float64,
		BuyingPrice:  dataDB.BuyingPrice.Float64,
		Uom1:         dataDB.Uom1.String,
		Uom2:         dataDB.Uom2.String,
		Conv1To2:     dataDB.Conv1To2.Int32,
		Division: dto.StructGeneral{
			ID:   dataDB.DivisionID.Int64,
			Code: dataDB.DivisionCode.String,
			Name: dataDB.DivisionName.String,
		},
		Category: dto.StructGeneral{
			ID:   dataDB.CategoryID.Int64,
			Code: dataDB.CategoryCode.String,
			Name: dataDB.CategoryName.String,
		},
		Group: dto.StructGeneral{
			ID:   dataDB.GroupID.Int64,
			Code: dataDB.GroupCode.String,
			Name: dataDB.GroupName.String,
		},
		CreatedAt: dataDB.CreatedAt,
		UpdatedAt: dataDB.UpdatedAt,
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorProduct(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "uq_prod_companyidcode") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
