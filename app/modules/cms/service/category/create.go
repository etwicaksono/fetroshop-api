package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *CategoryServiceImpl) Create(ctx *fiber.Ctx) (*appModel.Response, error) {
	svc.Logger.WebLoggerResetOutput()
	payload := new(model.UpsertCategoryRequest)
	err := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if err != nil {
		return nil, err
	}

	// TODO: implement me

	return &appModel.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Unimplemented", // #marked: message
	}, nil
}