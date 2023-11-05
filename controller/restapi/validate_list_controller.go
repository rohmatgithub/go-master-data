package restapi

import (
	"github.com/gofiber/fiber/v2"
	"go-master-data/dto"
	"go-master-data/model"
)

func validateList(c *fiber.Ctx, validOrderBy []string, validOperator map[string]dto.DefaultOperator) (dtoList dto.GetListRequest, listSearch []dto.SearchByParam, errMdl model.ErrorModel) {
	dtoList = dto.GetListRequest{
		Page:    c.QueryInt("page"),
		Limit:   c.QueryInt("limit"),
		OrderBy: c.Query("order_by"),
		Filter:  c.Query("filter"),
	}

	errMdl = dtoList.ValidateInputPageLimitAndOrderBy([]int{10, 30, 50, 100}, validOrderBy)
	if errMdl.Error != nil {
		return
	}

	listSearch, errMdl = dtoList.ValidateFilter(validOperator)
	return
}
