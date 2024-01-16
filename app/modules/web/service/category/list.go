package category

import (
	"fmt"

	ctEty "github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"gopkg.in/guregu/null.v3"
)

func (svc *CategoryServiceImpl) List(ctx *fiber.Ctx) (*appModel.Response, error) {
	var (
		categories []ctEty.Category
		parentID   null.Int
	)
	payload := new(model.ListCategoriesRequest)
	errorMap, err := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if err != nil {
		return responsehelper.Response500(constant.ERROR_GENERAL, fiber.Map{"message": err.Error()}), nil
	}
	if errorMap != nil {
		return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"messages": errorMap}), nil
	}

	if payload.ParentCode == "" {
		parentID = null.NewInt(0, false)
	} else {
		parent := new(ctEty.Category)
		result := svc.CategoryRepo.Find(parent, map[string]any{
			"code": payload.ParentCode,
		})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			svc.Logger.Error(result.Error.Error())
			return nil, errorhelper.Error500("Something went wrong") // #marked: message
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return nil, errorhelper.Error400("Invalid parent code") // #marked: message
		}
		parentID = null.NewInt(parent.ID, parent.ID != 0)
	}

	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)
	result := svc.CategoryRepo.List(&categories, fiber.Map{
		"parent_id": parentID,
	}, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.Error(result.Error.Error())
		return nil, errorhelper.Error500("Something went wrong") // #marked: message
	}

	var list []*model.CategoryResponse
	for _, ct := range categories {
		parentCode := ""
		if ct.Parent != nil {
			parentCode = ct.Parent.Code
		}

		category := &model.CategoryResponse{
			Code:         ct.Code,
			ParentCode:   null.NewString(parentCode, parentCode != ""),
			Name:         ct.Name,
			IsActive:     ct.IsActive,
			Icon:         ct.Icon,
			DisplayOrder: ct.DisplayOrder,
			CreatedAt:    ct.CreatedAt,
			UpdatedAt:    ct.UpdatedAt,
		}
		list = append(list, category)
	}

	return &appModel.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Successfuly got list of categories", // #marked: message
		Data:    list,
	}, nil
}
