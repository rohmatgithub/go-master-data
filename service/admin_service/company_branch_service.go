package admin_service

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/dto/admin_dto"
	"go-master-data/model"
)

type CompanyBranchService interface {
	Insert(request admin_dto.CompanyBranchRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Update(request admin_dto.CompanyBranchRequest, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Count(searchParam []dto.SearchByParam, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	ViewDetail(id int64, ctxModel *common.ContextModel) (dto.Payload, model.ErrorModel)
}
