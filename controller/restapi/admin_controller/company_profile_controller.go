package admin_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/model"
	"go-master-data/service/admin_service"
)

type CompanyProfileController struct {
	CompanyProfileService admin_service.CompanyProfileService
}

func NewCompanyProfileController(service admin_service.CompanyProfileService) CompanyProfileController {
	return CompanyProfileController{CompanyProfileService: service}
}

func (controller *CompanyProfileController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/companyprofile", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/companyprofile", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/companyprofile", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get(fmt.Sprintf("/companyprofile/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *CompanyProfileController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request dto.CompanyProfileRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyProfileService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyProfileController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request dto.CompanyProfileRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CompanyProfileService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyProfileController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyProfileService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CompanyProfileController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CompanyProfileService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
