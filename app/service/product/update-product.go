package product

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *productService) UpdateProduct(ctx *fiber.Ctx) (model.Response, error) {
	// TODO: implement me
	return responsehelper.Response500("not implemented", nil), nil
}
