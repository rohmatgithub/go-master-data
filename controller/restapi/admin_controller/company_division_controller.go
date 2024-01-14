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

type CompanyDivisionController struct {
	CompanyDivisionService admin_service.CompanyDivisionService
}

func NewCompanyDivisionController(service admin_service.CompanyDivisionService) CompanyDivisionController {
	return CompanyDivisionController{CompanyDivisionService: service}
}

func (controller *CompanyDivisionController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/companydivision", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/companydivision", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/companydivision", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get("/companydivision/initiate", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Count)
	})
	app.Get(fmt.Sprintf("/companydivision/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *CompanyDivisionController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request admin_dto.CompanyDivisionRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyDivisionService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyDivisionController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request admin_dto.CompanyDivisionRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyDivisionService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyDivisionController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyDivisionService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyDivisionController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name", "updated_at"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyDivisionService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyDivisionController) Count(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	listParam, errMdl := util_controller.ValidateCount(c, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyDivisionService.Count(listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
