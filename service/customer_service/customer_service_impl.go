package customer_service

import (
	"database/sql"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/dto/customer_dto"
	"go-master-data/entity"
	"go-master-data/entity/customer_entity"
	"go-master-data/model"
	"go-master-data/repository/customer_repository"
	"go-master-data/service"
	"strings"
	"time"
)

type customerServiceImpl struct {
	CustomerRepository customer_repository.CustomerRepository
}

func NewCustomerService(cpRepo customer_repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{CustomerRepository: cpRepo}
}

func (cp *customerServiceImpl) Insert(request customer_dto.CustomerRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(ctxModel)
	if validated != nil {
		out.Status.Detail = validated
		errMdl = model.GenerateFailedValidate()
		return
	}

	// check by code
	cpDb, errMdl := cp.CustomerRepository.FetchData(customer_entity.CustomerEntity{
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
	pEntity := customer_entity.CustomerEntity{
		AbstractEntity: entity.AbstractEntity{
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID:      sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.CompanyID, Valid: true},
		Code:           sql.NullString{String: request.Code, Valid: true},
		Name:           sql.NullString{String: request.Name, Valid: true},
		Phone:          sql.NullString{String: request.Phone, Valid: true},
		Email:          sql.NullString{String: request.Email, Valid: true},
		Address:        sql.NullString{String: request.Address, Valid: true},
		CountryID:      sql.NullInt64{Int64: request.CountryID, Valid: true},
		ProvinceID:     sql.NullInt64{Int64: request.ProvinceID, Valid: true},
		DistrictID:     sql.NullInt64{Int64: request.DistrictID, Valid: true},
		SubDistrictID:  sql.NullInt64{Int64: request.SubDistrictID, Valid: true},
		UrbanVillageID: sql.NullInt64{Int64: request.UrbanVillageID, Valid: true},
	}

	errMdl = cp.CustomerRepository.Insert(&pEntity)

	out.Status.Message = service.InsertI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *customerServiceImpl) Update(request customer_dto.CustomerRequest, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validated, errMdl := request.ValidateUpdate(ctxModel)
	if errMdl.Error != nil {
		return
	}
	if validated != nil {
		out.Status.Detail = validated
		return
	}
	cpDb, errMdl := cp.CustomerRepository.FetchData(customer_entity.CustomerEntity{
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
	pEntity := customer_entity.CustomerEntity{
		AbstractEntity: entity.AbstractEntity{
			ID:        request.ID,
			CreatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: ctxModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		CompanyID:      sql.NullInt64{Int64: ctxModel.AuthAccessTokenModel.CompanyID, Valid: true},
		Code:           sql.NullString{String: request.Code, Valid: true},
		Name:           sql.NullString{String: request.Name, Valid: true},
		Phone:          sql.NullString{String: request.Phone, Valid: true},
		Email:          sql.NullString{String: request.Email, Valid: true},
		Address:        sql.NullString{String: request.Address, Valid: true},
		CountryID:      sql.NullInt64{Int64: request.CountryID, Valid: true},
		ProvinceID:     sql.NullInt64{Int64: request.ProvinceID, Valid: true},
		DistrictID:     sql.NullInt64{Int64: request.DistrictID, Valid: true},
		SubDistrictID:  sql.NullInt64{Int64: request.SubDistrictID, Valid: true},
		UrbanVillageID: sql.NullInt64{Int64: request.UrbanVillageID, Valid: true},
	}

	errMdl = cp.CustomerRepository.Update(&pEntity)
	if errMdl.CausedBy != nil {
		errMdl = convertErrorProduct(errMdl)
	}
	out.Status.Message = service.UpdateI18NMessage(ctxModel.AuthAccessTokenModel.Locale)

	return
}

func (cp *customerServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := cp.CustomerRepository.List(dtoList, searchParam, ctxModel)
	if errMdl.Error != nil {
		return
	}

	var result []customer_dto.ListCustomerResponse
	for _, temp := range resultDB {
		data := temp.(customer_entity.CustomerEntity)
		result = append(result, customer_dto.ListCustomerResponse{
			ID:        data.ID,
			Code:      data.Code.String,
			Name:      data.Name.String,
			Phone:     data.Phone.String,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	out.Data = result
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *customerServiceImpl) Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	resultDB, errMdl := cp.CustomerRepository.Count(searchParam, ctxModel)
	if errMdl.Error != nil {
		return
	}
	out.Data = resultDB
	out.Status.Message = service.ListI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *customerServiceImpl) ViewDetail(id int64, ctxModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	if id < 1 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}

	dataDB, errMdl := cp.CustomerRepository.View(id)
	if errMdl.Error != nil {
		return
	}

	if dataDB.ID == 0 {
		errMdl = model.GenerateUnknownDataError(constanta.ID)
		return
	}
	out.Data = customer_dto.CustomerDetailResponse{
		ID:        dataDB.ID,
		Code:      dataDB.Code.String,
		Name:      dataDB.Name.String,
		Phone:     dataDB.Phone.String,
		Email:     dataDB.Email.String,
		Address:   dataDB.Address.String,
		CreatedAt: dataDB.CreatedAt,
		UpdatedAt: dataDB.UpdatedAt,
		Country: dto.StructGeneral{
			ID:   dataDB.CountryID.Int64,
			Code: dataDB.CountryCode.String,
			Name: dataDB.CountryName.String,
		},
		Province: dto.StructGeneral{
			ID:   dataDB.ProvinceID.Int64,
			Code: dataDB.ProvinceCode.String,
			Name: dataDB.ProvinceName.String,
		},
		District: dto.StructGeneral{
			ID:   dataDB.DistrictID.Int64,
			Code: dataDB.DistrictCode.String,
			Name: dataDB.DistrictName.String,
		},
		SubDistrict: dto.StructGeneral{
			ID:   dataDB.SubDistrictID.Int64,
			Code: dataDB.SubDistrictCode.String,
			Name: dataDB.SubDistrictName.String,
		},
		UrbanVillage: dto.StructGeneral{
			ID:   dataDB.UrbanVillageID.Int64,
			Code: dataDB.UrbanVillageCode.String,
			Name: dataDB.UrbanVillageName.String,
		},
	}

	out.Status.Message = service.ViewI18NMessage(ctxModel.AuthAccessTokenModel.Locale)
	return
}

func convertErrorProduct(err model.ErrorModel) (errMdl model.ErrorModel) {
	if err.CausedBy == nil {
		return err
	}
	if strings.Contains(err.CausedBy.Error(), "uq_cust_companyidbranchidcode") {
		return model.GenerateHasUsedDataError(constanta.Code)
	}
	return
}
