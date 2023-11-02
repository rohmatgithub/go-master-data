package restapi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"go-master-data/common"
	"go-master-data/config"
	"go-master-data/constanta"
	"go-master-data/dto"
	"go-master-data/model"
	"time"
)

type abstractController struct {
}

func (ae abstractController) EndpointClientCredentials(c *fiber.Ctx, runFunc func(*fiber.Ctx, *common.ContextModel) (dto.Payload, model.ErrorModel)) error {
	// validate client_id
	tokenStr := c.Get(constanta.TokenHeaderNameConstanta)

	validateFunc := func(contextModel *common.ContextModel) (errMdl model.ErrorModel) {
		// cek token expired
		_, errMdl = model.JWTToken{}.ParsingJwtTokenInternal(tokenStr)
		if errMdl.Error != nil {
			return
		}

		return
	}

	return ae.serve(c, validateFunc, runFunc)
}

func (ae abstractController) serve(c *fiber.Ctx,
	validateFunc func(contextModel *common.ContextModel) model.ErrorModel,
	runFunc func(*fiber.Ctx, *common.ContextModel) (dto.Payload, model.ErrorModel)) (err error) {
	var (
		response     dto.StandardResponse
		payload      dto.Payload
		contextModel common.ContextModel
	)

	requestID := c.Locals("requestid").(string)
	logModel := c.Context().Value(constanta.ApplicationContextConstanta).(*common.LoggerModel)

	contextModel.LoggerModel = *logModel
	response.Header = dto.HeaderResponse{
		RequestID: requestID,
		Version:   config.ApplicationConfiguration.GetServerConfig().Version,
		Timestamp: time.Now().Format(time.RFC3339),
	}

	defer func() {
		response.Payload = payload

		adaptor.CopyContextToFiberContext(logModel, c.Context())
		err = c.JSON(response)
	}()
	// validate
	errMdl := validateFunc(&contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, logModel, &payload, errMdl)
		return
	}
	payload, errMdl = runFunc(c, &contextModel)
	if errMdl.Error != nil {
		generateEResponseError(c, logModel, &payload, errMdl)
	} else {
		payload.Status = dto.StatusPayload{
			Success: true,
			Code:    "OK",
		}
	}
	return
}

func generateEResponseError(c *fiber.Ctx, logModel *common.LoggerModel, payload *dto.Payload, errMdl model.ErrorModel) {
	logModel.Code = errMdl.Error.Error()
	logModel.Class = errMdl.Line
	if errMdl.CausedBy != nil {
		logModel.Message = errMdl.CausedBy.Error()
	}
	// write failed
	c.Status(errMdl.Code)
	payload.Status = dto.StatusPayload{
		Success: false,
		Code:    errMdl.Error.Error(),
		Message: "",
	}
}
