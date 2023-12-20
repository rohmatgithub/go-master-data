package admin_service

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
)

type CompanyProfileService interface {
	Insert(request dto.CompanyProfileRequest, contextModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	Update(request dto.CompanyProfileRequest, contextModel *common.ContextModel) (dto.Payload, model.ErrorModel)
	List(dtoList dto.GetListRequest, searchParam []dto.SearchByParam) (dto.Payload, model.ErrorModel)
	ViewDetail(id int64, contextModel *common.ContextModel) (dto.Payload, model.ErrorModel)
}
