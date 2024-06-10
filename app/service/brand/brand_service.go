package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BrandService interface {
	CreateBrand(ctx *fiber.Ctx) (model.Response, error)
	UpdateBrand(ctx *fiber.Ctx) (model.Response, error)
	DeleteBrand(ctx *fiber.Ctx) (model.Response, error)
	ListBrand(ctx *fiber.Ctx) (model.Response, error)
	ListBrandByPrefix(ctx *fiber.Ctx) (model.Response, error)
	FindBrand(ctx *fiber.Ctx) (model.Response, error)
}

type brandService struct {
	Err       error
	DB        *gorm.DB
	Config    *viper.Viper
	Validate  *validator.Validate
	Logger    *logger.Logger
	BrandRepo brands.BrandRepo
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	brandRepo brands.BrandRepo,
) BrandService {
	return &brandService{
		Err:       conn.Err,
		DB:        conn.DB,
		Config:    config,
		Validate:  validate,
		Logger:    logger,
		BrandRepo: brandRepo,
	}
}
