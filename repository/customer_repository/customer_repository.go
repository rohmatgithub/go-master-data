package customer_repository

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/entity/customer_entity"
	"go-master-data/model"
)

type CustomerRepository interface {
	Insert(entity *customer_entity.CustomerEntity) model.ErrorModel
	Update(entity *customer_entity.CustomerEntity) model.ErrorModel
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result []interface{}, errMdl model.ErrorModel)
	View(id int64) (customer_entity.CustomerDetailEntity, model.ErrorModel)
	FetchData(entity customer_entity.CustomerEntity) (customer_entity.CustomerEntity, model.ErrorModel)
	Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (result int64, errMdl model.ErrorModel)
}
