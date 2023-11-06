package admin_controller

import (
	"github.com/gofiber/fiber/v2"
	"go-master-data/common"
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
