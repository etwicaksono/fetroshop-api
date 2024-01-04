package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (rg *RegistrationServiceImpl) Register(request *webModel.RegistrationRequest) (*model.Response, error) {
	// TODO - check if user already exists
	// TODO - insert user into database
	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: "Not Implemented",
		Data:    request,
	}, nil
}
