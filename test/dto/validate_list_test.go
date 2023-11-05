package dto

import (
	"github.com/stretchr/testify/assert"
	"go-master-data/dto"
	"testing"
)

func TestValidateFilter(t *testing.T) {
	dto.GenerateValidOperator()
	dtoIn := dto.GetListRequest{
		Filter: "name like test, code like test",
	}

	searchParam, errMdl := dtoIn.ValidateFilter(dto.ValidOperatorRegional)
	assert.Nil(t, errMdl.Error)
	assert.NotZero(t, searchParam)
	assert.EqualValues(t, len(searchParam), 2)

	dtoIn = dto.GetListRequest{
		Filter: "names like test",
	}
	searchParam, errMdl = dtoIn.ValidateFilter(dto.ValidOperatorRegional)
	assert.NotNil(t, errMdl.Error)
}
