package product

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/helper/miniohelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(ctx *fiber.Ctx) (model.Response, error)
	UpdateProduct(ctx *fiber.Ctx) (model.Response, error)
	DeleteProduct(ctx *fiber.Ctx) (model.Response, error)
	ListProduct(ctx *fiber.Ctx) (model.Response, error)
	FindProduct(ctx *fiber.Ctx) (model.Response, error)
}

type productService struct {
	Err      error
	DB       *gorm.DB
	Config   *viper.Viper
	Validate *validator.Validate
	Logger   *logger.Logger
	Minio    *miniohelper.Minio
}

func ServiceProvider(
	conn *connection.Connection,
	config *viper.Viper,
	validate *validator.Validate,
	logger *logger.Logger,
	myMinio *miniohelper.Minio,
) ProductService {
	return &productService{
		Err:      conn.Err,
		DB:       conn.DB,
		Config:   config,
		Validate: validate,
		Logger:   logger,
		Minio:    myMinio,
	}
}
