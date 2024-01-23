package __file_test_admin

import (
	"github.com/stretchr/testify/assert"
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/dto/admin_dto"
	"go-master-data/service/admin_service"
	"go-master-data/test/repository/admin_repository"
	"testing"
)

func prepareCommon() {
	common.Validation = common.NewGoValidator()
}
func TestCompanyProfileService_Insert(t *testing.T) {
	// prepare config
	common.Validation = common.NewGoValidator()

	cpDao := admin_repository.NewCompanyProfileRepository()
	cpService := admin_service.NewCompanyProfileService(cpDao)

	// create input data request
	cpDto := admin_dto.CompanyProfileRequest{
		NPWP:           "NPWPTEST",
		Name:           "",
		Address1:       "Jakarta",
		Address2:       "Jakarta",
		CountryID:      1,
		ProvinceID:     1,
		DistrictID:     1,
		SubDistrictID:  1,
		UrbanVillageID: 1,
		AbstractDto:    dto.AbstractDto{},
	}

	// execute function
	resultDto, errMdl := cpService.Insert(cpDto, &common.ContextModel{})
	assert.NotNil(t, errMdl.Error)
	assert.NotNil(t, resultDto.Status.Detail)
	mapDetail := resultDto.Status.Detail.(map[string]string)
	assert.Equal(t, "must be a valid numeric value", mapDetail["npwp"])
}
