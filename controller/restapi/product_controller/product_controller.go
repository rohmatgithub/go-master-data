package product_controller

import (
	"fmt"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/dto/product_dto"
	"go-master-data/model"
	"go-master-data/service/product_service"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService product_service.ProductService
}

func NewProductController(service product_service.ProductService) ProductController {
	return ProductController{ProductService: service}
}

func (controller *ProductController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/product", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/product", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/product", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get("/product/count", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Count)
	})
	app.Get(fmt.Sprintf("/product/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *ProductController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request product_dto.ProductRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ProductService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request product_dto.ProductRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ProductService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name", "updated_at"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductController) Count(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	listParam, errMdl := util_controller.ValidateCount(c, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductService.Count(listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
