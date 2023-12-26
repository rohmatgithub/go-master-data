package admin_repository

import (
	"go-master-data/dto"
	"go-master-data/entity/admin_entity"
	"go-master-data/model"
)

type CompanyBranchRepository interface {
	Insert(entity *admin_entity.CompanyBranchEntity) model.ErrorModel
	Update(entity *admin_entity.CompanyBranchEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (result []interface{}, errMdl model.ErrorModel)
	View(id int64) (admin_entity.CompanyBranchDetailEntity, model.ErrorModel)
	FetchData(entity admin_entity.CompanyBranchEntity) (admin_entity.CompanyBranchEntity, model.ErrorModel)
}
