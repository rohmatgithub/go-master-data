package admin_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
)

type CompanyDivisionRepository interface {
	Insert(entity *admin_entity.CompanyDivisionEntity) model.ErrorModel
	Update(entity *admin_entity.CompanyDivisionEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
	Count(searchParam []dto.SearchByParam) (result int64, errMdl model.ErrorModel)
	View(id int64) (admin_entity.CompanyDivisionDetailEntity, model.ErrorModel)
	FetchData(entity admin_entity.CompanyDivisionEntity) (admin_entity.CompanyDivisionEntity, model.ErrorModel)
}
