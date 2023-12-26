package product_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/dto/product_dto"
	"go-master-data/model"
	"go-master-data/service/product_service"
)

type ProductGroupController struct {
	ProductGroupService product_service.ProductGroupService
}

func NewProductGroupController(service product_service.ProductGroupService) ProductGroupController {
	return ProductGroupController{ProductGroupService: service}
}

func (controller *ProductGroupController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/productgroup", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/productgroup", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/productgroup", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get(fmt.Sprintf("/productgroup/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *ProductGroupController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request product_dto.ProductGroupRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ProductGroupService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductGroupController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request product_dto.ProductGroupRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ProductGroupService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductGroupController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductGroupService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductGroupController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductGroupService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
