package admin_controller

import (
	"fmt"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/dto/admin_dto"
	"go-master-data/model"
	"go-master-data/service/admin_service"

	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	CompanyService admin_service.CompanyService
}

func NewCompanyController(service admin_service.CompanyService) CompanyController {
	return CompanyController{CompanyService: service}
}

func (controller *CompanyController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/company", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/company", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/company", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get("/company/initiate", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Count)
	})
	app.Get(fmt.Sprintf("/company/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *CompanyController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request admin_dto.CompanyRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request admin_dto.CompanyRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name", "updated_at"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyController) Count(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	listParam, errMdl := util_controller.ValidateCount(c, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyService.Count(listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
