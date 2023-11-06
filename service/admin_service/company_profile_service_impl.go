package admin_service

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/model"
)

type companyProfileImpl struct {
}

func NewCompanyProfileService() CompanyProfileService {
	return &companyProfileImpl{}
}

func (cp *companyProfileImpl) Insert(request dto.CompanyProfileRequest, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {

	out.Status.Detail = request.ValidateInsert(contextModel)

	return
}
