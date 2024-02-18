package __file_test_admin

import (
	"go-master-data/common"
	"go-master-data/dto"
	"go-master-data/dto/admin_dto"
	"go-master-data/service/admin_service"
	"go-master-data/test/repository/admin_repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepareCommon() {
	common.Validation = common.NewGoValidator()
}
func TestCompanyProfileService_Insert_Failed(t *testing.T) {
	// prepare config
	common.Validation = common.NewGoValidator()

	cpDao := admin_repository.NewCompanyProfileRepository()
	cpService := admin_service.NewCompanyProfileService(cpDao)

	// create input data request
	cpDto := admin_dto.CompanyProfileRequest{
		NPWP:           "883923",
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

func TestCompanyProfileService_Insert_Success(t *testing.T) {
	// prepare config
	common.Validation = common.NewGoValidator()

	cpDao := admin_repository.NewCompanyProfileRepository()
	cpService := admin_service.NewCompanyProfileService(cpDao)

	// create input data request
	cpDto := admin_dto.CompanyProfileRequest{
		NPWP:           "8839230989808768",
		Name:           "Sinde Budi",
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
	assert.Nil(t, errMdl.Error)
	assert.Nil(t, resultDto.Status.Detail)
}
