package regional_controller

import (
	"github.com/gofiber/fiber/v2"
	"go-master-data/common"
	"go-master-data/controller/restapi/util_controller"
	"go-master-data/dto"
	"go-master-data/model"
	"go-master-data/service/regional_service"
)

type RegionalController struct {
	DistrictService    regional_service.DistrictService
	SubDistrictService regional_service.SubDistrictService
	UrbanVillage       regional_service.UrbanVillageService
}

func NewRegionalController(
	district regional_service.DistrictService,
	subDistrict regional_service.SubDistrictService,
	urbanVillage regional_service.UrbanVillageService) RegionalController {
	return RegionalController{
		DistrictService:    district,
		SubDistrictService: subDistrict,
		UrbanVillage:       urbanVillage,
	}
}
func (controller *RegionalController) Route(app fiber.Router) {
	var ae util_controller.AbstractController
	app.Get("/district", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.ListDistrict)
	})

	app.Get("/subdistrict", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.ListSubDistrict)
	})

	app.Get("/urbanvillage", func(c *fiber.Ctx) error {
		return ae.ServeJwtToken(c, "", controller.ListUrbanVillage)
	})
}

func (controller *RegionalController) ListDistrict(c *fiber.Ctx, _ *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorRegional)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.DistrictService.List(dtoList, listParam)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *RegionalController) ListSubDistrict(c *fiber.Ctx, _ *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorRegional)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.SubDistrictService.List(dtoList, listParam)
	if errMdl.Error != nil {
		return
	}

	return
}

func (controller *RegionalController) ListUrbanVillage(c *fiber.Ctx, _ *common.ContextModel) (out dto.Payload, errMdl model.ErrorModel) {
	// set to search param
	dtoList, listParam, errMdl := util_controller.ValidateList(c, []string{"id", "code", "name"}, dto.ValidOperatorRegional)
	if errMdl.Error != nil {
		return
	}
	out, errMdl = controller.UrbanVillage.List(dtoList, listParam)
	if errMdl.Error != nil {
		return
	}

	return
}
