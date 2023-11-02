package restapi

import (
	"github.com/gofiber/fiber/v2"
	"go-master-data/service/regional_service"
)

type regionalController struct {
	SubDistrictService regional_service.SubDistrictService
}

func (controller regionalController) Route(app fiber.Router) {
	//var ae abstractController
	//app.Post("/verify", func(ctx *fiber.Ctx) error {
	//	return ae.EndpointClientCredentials(ctx, credentialsservice.CredentialsService.VerifyService)
	//})
}
