package brand

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *brandService) List(ctx *fiber.Ctx) (model.Response, error) {
	var brandSlice []brands.Brand
	payload := new(model.ListBrandsRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return model.Response{}, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)
	result := svc.BrandRepo.List(&brandSlice, fiber.Map{}, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return model.Response{}, result.Error
	}

	var list []*model.BrandResponse
	for _, brd := range brandSlice {
		brand := &model.BrandResponse{
			Code:      brd.Code,
			Name:      brd.Name,
			IsActive:  brd.IsActive,
			Icon:      brd.Icon,
			CreatedAt: brd.CreatedAt,
			UpdatedAt: brd.UpdatedAt,
		}
		list = append(list, brand)
	}

	return responsehelper.Response200(
		"Successfuly got list of brands", // #marked: message
		list,
		fiber.Map{"total": result.RowsAffected},
	), nil
}
