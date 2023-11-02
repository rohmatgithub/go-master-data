package restapi

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go-master-data/common"
	"go-master-data/config"
	"go-master-data/constanta"
	"os"
	"strconv"
)

func middleware(c *fiber.Ctx) error {
	logModel := &common.LoggerModel{
		Pid:         strconv.Itoa(os.Getpid()),
		RequestID:   c.Locals("requestid").(string),
		Resource:    "",
		Application: config.ApplicationConfiguration.GetServerConfig().ResourceID,
		Version:     config.ApplicationConfiguration.GetServerConfig().Version,
		ByteIn:      len(c.Body()),
		//Path:        c.BaseURL(),
	}
	logger := context.WithValue(c.Context(), constanta.ApplicationContextConstanta, logModel)
	adaptor.CopyContextToFiberContext(logger, c.Context())

	err := c.Next()
	if err != nil {
		return err
	}
	logModel = c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)
	logModel.Status = c.Response().StatusCode()
	logModel.Path = c.OriginalURL()
	log.Info(common.GenerateLogModel(*logModel))
	return err
}

func NotFoundHandler(c *fiber.Ctx) error {
	// Customize the response for the 404 error
	return c.Status(fiber.StatusNotFound).SendString("404 Not Found")
}

func customErrorHandler(c *fiber.Ctx, err error) {
	// Handle the error here
	fmt.Printf("Error: %v\n", err)

	// Return a custom error response
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Something went wrong",
	})
}
