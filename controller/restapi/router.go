package restapi

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-master-data/common"
	"go-master-data/config"
	"go-master-data/controller/restapi/admin_controller"
	"go-master-data/controller/restapi/example_controller"
	"go-master-data/controller/restapi/regional_controller"
	"go-master-data/repository/admin_repository"
	"go-master-data/repository/example_repository"
	"go-master-data/repository/regional_repository"
	"go-master-data/service/admin_service"
	"go-master-data/service/example_service"
	"go-master-data/service/regional_service"
)

func Router() error {
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Oauth App v1.0.0",
		ColorScheme:   fiber.Colors{Green: ""},
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})
	app.Use(requestid.New())
	//app.Use(recoverfiber.New())
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				customErrorHandler(c, fmt.Errorf("%v", r))
			}
		}()
		return c.Next()
	})
	app.Use(middleware)
	//file, err := os.OpenFile("fiber.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//iw := io.MultiWriter(os.Stdout, file)
	//defer file.Close()
	//app.Use(logger.New(logger.Config{
	//	Format:     "[${time}] pid:${pid}, request-id:${locals:requestid}, status:${status}, method:${method}, path:${path}, error-message:[${error}]\n",
	//	TimeFormat: time.DateTime,
	//	TimeZone:   "Asia/Jakarta",
	//	Output:     iw,
	//}))

	v1 := app.Group("/v1/master")

	exampleRepository := example_repository.NewExampleRepository(common.GormDB)
	exampleService := example_service.NewExampleService(exampleRepository)
	exampleController := example_controller.NewExampleController(exampleService)
	exampleController.Route(v1)

	districtRepository := regional_repository.NewDistrictRepository(common.GormDB)
	districtService := regional_service.NewDistrictService(districtRepository)

	subDistrictRepository := regional_repository.NewSubDistrictRepository(common.GormDB)
	subDistrictService := regional_service.NewSubDistrictService(districtRepository, subDistrictRepository)

	urbanVillageRepository := regional_repository.NewUrbanVillageRepository(common.GormDB)
	urbanVillageService := regional_service.NewUrbanVillageService(subDistrictRepository, urbanVillageRepository)

	regionalController := regional_controller.NewRegionalController(districtService, subDistrictService, urbanVillageService)
	regionalController.Route(v1)

	companyProfileRepository := admin_repository.NewCompanyProfileRepository(common.GormDB)
	companyProfileService := admin_service.NewCompanyProfileService(companyProfileRepository)
	companyProfileController := admin_controller.NewCompanyProfileController(companyProfileService)
	companyProfileController.Route(v1)

	app.Use(NotFoundHandler)
	return app.Listen(fmt.Sprintf(":%d", config.ApplicationConfiguration.GetServerConfig().Port))
}
