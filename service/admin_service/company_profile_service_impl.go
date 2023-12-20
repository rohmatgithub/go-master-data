package admin_service

import (
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/entity"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
	"go-master-data/repository/admin_repository"
	"go-master-data/service"
	"time"
)

type companyProfileServiceImpl struct {
	CompanyProfileRepository admin_repository.CompanyProfileRepository
}

func NewCompanyProfileService(cpRepo admin_repository.CompanyProfileRepository) CompanyProfileService {
	return &companyProfileServiceImpl{CompanyProfileRepository: cpRepo}
}

func (cp *companyProfileServiceImpl) Insert(request dto.CompanyProfileRequest, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	validated := request.ValidateInsert(contextModel)
	if validated != nil {
		out.Status.Detail = validated
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
			CreatedBy: contextModel.AuthAccessTokenModel.ResourceUserID,
			UpdatedBy: contextModel.AuthAccessTokenModel.ResourceUserID,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		NPWP:           request.NPWP,
		Name:           request.Name,
		Address1:       request.Address1,
		Address2:       request.Address2,
		CountryID:      request.CountryID,
		DistrictID:     request.DistrictID,
		SubDistrictID:  request.SubDistrictID,
		UrbanVillageID: request.UrbanVillageID,
	}

	errMdl = cp.CompanyProfileRepository.Insert(&cpEntity)

	out.Status.Message = service.InsertI18NMessage(contextModel.AuthAccessTokenModel.Locale)
	return
}

func (cp *companyProfileServiceImpl) Update(request dto.CompanyProfileRequest, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	validate, errMdl := request.ValidateUpdate(contextModel)
	if errMdl.Error != nil {

	}
	return
}

func (cp *companyProfileServiceImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (out dto.Payload, errMdl model.ErrorModel) {
	return
}

func (cp *companyProfileServiceImpl) ViewDetail(id int64, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	return
}
