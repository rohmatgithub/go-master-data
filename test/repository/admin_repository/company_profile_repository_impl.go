package admin_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
	"go-master-data/repository/admin_repository"
)

type companyProfileRepositoryImpl struct {
}

func NewCompanyProfileRepository() admin_repository.CompanyProfileRepository {
	return &companyProfileRepositoryImpl{}
}

func (repo *companyProfileRepositoryImpl) Insert(cp *admin_entity.CompanyProfileEntity) (errMdl model.ErrorModel) {

	return
}

func (repo *companyProfileRepositoryImpl) Update(cp *admin_entity.CompanyProfileEntity) (errMdl model.ErrorModel) {
	return
}

func (repo *companyProfileRepositoryImpl) List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel) {
	return

}

func (repo *companyProfileRepositoryImpl) Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel) {
	return
}

func (repo *companyProfileRepositoryImpl) View(id int64) (result admin_entity.CompanyProfileDetailEntity, errMdl model.ErrorModel) {
	return
}

func (repo *companyProfileRepositoryImpl) FetchData(entity admin_entity.CompanyProfileEntity) (result admin_entity.CompanyProfileEntity, errMdl model.ErrorModel) {
	return
}
