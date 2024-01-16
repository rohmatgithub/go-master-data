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

type CompanyBranchController struct {
	CompanyBranchService admin_service.CompanyBranchService
}

func NewCompanyBranchController(service admin_service.CompanyBranchService) CompanyBranchController {
	return CompanyBranchController{CompanyBranchService: service}
}

func (controller *CompanyBranchController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/companybranch", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/companybranch", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/companybranch", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get("/companybranch/initiate", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Count)
	})
	app.Get(fmt.Sprintf("/companybranch/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *CompanyBranchController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request admin_dto.CompanyBranchRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyBranchService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyBranchController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request admin_dto.CompanyBranchRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyBranchService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyBranchController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyBranchService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyBranchController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, dto.DefaultOrder, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyBranchService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyBranchController) Count(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	listParam, errMdl := util_controller.ValidateCount(c, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyBranchService.Count(listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
