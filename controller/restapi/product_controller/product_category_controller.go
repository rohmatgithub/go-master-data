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

type ProductCategoryController struct {
	ProductCategoryService product_service.ProductCategoryService
}

func NewProductCategoryController(service product_service.ProductCategoryService) ProductCategoryController {
	return ProductCategoryController{ProductCategoryService: service}
}

func (controller *ProductCategoryController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/productcategory", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/productcategory", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/productcategory", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get("/productcategory/count", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Count)
	})
	app.Get(fmt.Sprintf("/productcategory/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *ProductCategoryController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request product_dto.ProductCategoryRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ProductCategoryService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductCategoryController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request product_dto.ProductCategoryRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ProductCategoryService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductCategoryController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductCategoryService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductCategoryController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductCategoryService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ProductCategoryController) Count(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	listParam, errMdl := util_controller.ValidateCount(c, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ProductCategoryService.Count(listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
