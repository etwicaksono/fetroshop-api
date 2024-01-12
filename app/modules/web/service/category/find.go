package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *CategoryServiceImpl) Find(ctx *fiber.Ctx) (*appModel.Response, error) {
	payload := new(model.FindCategoryRequest)
	validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	// TODO: implement this
	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Unimplemented", // #marked: message
		Data:    payload,
	}, nil
}