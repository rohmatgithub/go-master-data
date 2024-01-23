package customer_controller

import (
	"fmt"
	"go-master-data/common"
	"go-master-data/constanta"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/dto/customer_dto"
	"go-master-data/model"
	"go-master-data/service/customer_service"

	"github.com/gofiber/fiber/v2"
)

type CustomerController struct {
	CustomerService customer_service.CustomerService
}

func NewCustomerController(service customer_service.CustomerService) CustomerController {
	return CustomerController{CustomerService: service}
}

func (controller *CustomerController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Post("/customer", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Insert)
	})
	app.Put("/customer", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Update)
	})
	app.Get("/customer", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.List)
	})
	app.Get("/customer/count", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.Count)
	})
	app.Get(fmt.Sprintf("/customer/:%s", constanta.ParamID), func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.View)
	})
}
func (controller *CustomerController) Insert(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request customer_dto.CustomerRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CustomerService.Insert(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CustomerController) Update(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	var request customer_dto.CustomerRequest
	err := c.BodyParser(&request)
	if err != nil {
		errMdl = model.GenerateInvalidRequestError(err)
		return
	}
	out, errMdl = controller.CustomerService.Update(request, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CustomerController) View(c *fiber.Ctx, contextModel *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	id, errMdl := util_controller.GetParamID(c)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CustomerService.ViewDetail(id, contextModel)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CustomerController) List(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name", "updated_at"}, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CustomerService.List(dtoList, listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *CustomerController) Count(c *fiber.Ctx, ctx *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	listParam, errMdl := util_controller.ValidateCount(c, dto.ValidOperatorGeneral)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.CustomerService.Count(listParam, ctx)
	if errMdl.Error != nil {
		return
	}

	return
}
