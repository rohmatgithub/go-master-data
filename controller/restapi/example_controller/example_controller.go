package example_controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/model"
	"go-master-data/service/example_service"
)

type ExampleController struct {
	ExampleService example_service.ExampleService
}

func NewExampleController(service example_service.ExampleService) ExampleController {
	return ExampleController{ExampleService: service}
}

func (controller *ExampleController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/example", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/example", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/example", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get(fmt.Sprintf("/example/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *ExampleController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request dto.ExampleRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ExampleService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ExampleController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request dto.ExampleRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.ExampleService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ExampleController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ExampleService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *ExampleController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.ExampleService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
