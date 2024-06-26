package responsehelper

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func Response200(msg string, data interface{}, meta interface{}) model.Response {
	return model.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Response201(msg string, data interface{}, meta interface{}) model.Response {
	return model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: msg,
		Data:    data,
		Meta:    meta,
	}
}

func Response400(msg string, meta interface{}) model.Response {
	return model.Response{
		Code:    fiber.StatusBadRequest,
		Status:  utils.StatusMessage(fiber.StatusBadRequest),
		Message: msg,
		Meta:    meta,
	}
}

func Response401(msg string, meta interface{}) model.Response {
	return model.Response{
		Code:    fiber.StatusUnauthorized,
		Status:  utils.StatusMessage(fiber.StatusUnauthorized),
		Message: msg,
		Meta:    meta,
	}
}

func Response500(msg string, meta interface{}) model.Response {
	return model.Response{
		Code:    fiber.StatusInternalServerError,
		Status:  utils.StatusMessage(fiber.StatusInternalServerError),
		Message: msg,
		Meta:    meta,
	}
}

func ResponseErrorValidation(errValidation fiber.Map) model.Response {
	return model.Response{
		Code:    fiber.StatusBadRequest,
		Status:  utils.StatusMessage(fiber.StatusBadRequest),
		Message: constant.ERROR_VALIDATION,
		Errors:  errValidation,
	}
}
