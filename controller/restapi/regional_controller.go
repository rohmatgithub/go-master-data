package restapi

import (
	"github.com/gofiber/fiber/v2"
	"go-master-data/dto"
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
	app.Get("/district", controller.ListDistrict)
}

func (controller *RegionalController) ListDistrict(c *fiber.Ctx) error {
	response, errMdl := controller.DistrictService.List(dto.GetListRequest{}, nil)
	if errMdl.Error != nil {
		return errMdl.Error
	}

	return c.JSON(response)
}
