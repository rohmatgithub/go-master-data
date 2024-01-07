package admin_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
)

type CompanyProfileRepository interface {
	Insert(entity *admin_entity.CompanyProfileEntity) model.ErrorModel
	Update(entity *admin_entity.CompanyProfileEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
	Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel)
	View(id int64) (admin_entity.CompanyProfileDetailEntity, model.ErrorModel)
	FetchData(entity admin_entity.CompanyProfileEntity) (admin_entity.CompanyProfileEntity, model.ErrorModel)
}
