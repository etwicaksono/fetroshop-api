// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	postgres5 "github.com/BuildWithYou/fetroshop-api/app/domain/categories/postgres"
	postgres2 "github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses/postgres"
	postgres4 "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses/postgres"
	postgres3 "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	controller2 "github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/service/category"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeWebServer() *app.Fetroshop {
	viper := confighelper.GetConfig()
	docsDocs := docs.DocsProvider(viper)
	connectionDBType := _wireDBTypeValue
	loggerLogger := logger.NewWebLogger(viper)
	connectionConnection := connection.OpenDBConnection(connectionDBType, viper, loggerLogger)
	userAccessRepo := postgres.RepoProvider(connectionConnection)
	customerAccessRepo := postgres2.RepoProvider(connectionConnection)
	jwtMiddleware := middleware.JwtMiddlewareProvider(viper, userAccessRepo, customerAccessRepo)
	dbMiddleware := middleware.DBMiddlewareProvider(connectionConnection, loggerLogger)
	validate := validatorhelper.GetValidator()
	userRepo := postgres3.RepoProvider(connectionConnection)
	customerRepo := postgres4.RepoProvider(connectionConnection)
	authService := auth.ServiceProvider(connectionConnection, viper, validate, userRepo, userAccessRepo, customerRepo, customerAccessRepo, loggerLogger)
	authController := controller.AuthControllerProvider(validate, authService)
	categoryRepo := postgres5.RepoProvider(connectionConnection)
	categoryService := category.ServiceProvider(connectionConnection, viper, validate, loggerLogger, categoryRepo)
	categoryController := controller.CategoryControllerProvider(validate, categoryService)
	controllerController := controller.WebControllerProvider(authController, categoryController)
	router := web.RouterProvider(docsDocs, jwtMiddleware, dbMiddleware, controllerController, loggerLogger)
	serverConfig := web.WebServerConfigProvider(router, loggerLogger)
	fetroshop := app.CreateFiber(serverConfig)
	return fetroshop
}

var (
	_wireDBTypeValue = dbType
)

func InitializeCmsServer() *app.Fetroshop {
	viper := confighelper.GetConfig()
	docsDocs := docs.DocsProvider(viper)
	connectionDBType := _wireDBTypeValue
	loggerLogger := logger.NewCmsLogger(viper)
	connectionConnection := connection.OpenDBConnection(connectionDBType, viper, loggerLogger)
	userAccessRepo := postgres.RepoProvider(connectionConnection)
	customerAccessRepo := postgres2.RepoProvider(connectionConnection)
	jwtMiddleware := middleware.JwtMiddlewareProvider(viper, userAccessRepo, customerAccessRepo)
	dbMiddleware := middleware.DBMiddlewareProvider(connectionConnection, loggerLogger)
	validate := validatorhelper.GetValidator()
	userRepo := postgres3.RepoProvider(connectionConnection)
	customerRepo := postgres4.RepoProvider(connectionConnection)
	authService := auth.ServiceProvider(connectionConnection, viper, validate, userRepo, userAccessRepo, customerRepo, customerAccessRepo, loggerLogger)
	authController := controller2.AuthControllerProvider(authService)
	categoryRepo := postgres5.RepoProvider(connectionConnection)
	categoryService := category.ServiceProvider(connectionConnection, viper, validate, loggerLogger, categoryRepo)
	categoryController := controller2.CategoryControllerProvider(validate, categoryService)
	controllerController := controller2.CmsControllerProvider(authController, categoryController)
	router := cms.RouterProvider(docsDocs, jwtMiddleware, dbMiddleware, controllerController, loggerLogger)
	serverConfig := cms.CmsServerConfigProvider(router, loggerLogger)
	fetroshop := app.CreateFiber(serverConfig)
	return fetroshop
}

// injector.go:

var dbType connection.DBType = connection.DB_MAIN

var repoSet = wire.NewSet(postgres4.RepoProvider, postgres2.RepoProvider, postgres3.RepoProvider, postgres.RepoProvider, postgres5.RepoProvider)

var serviceSet = wire.NewSet(auth.ServiceProvider, category.ServiceProvider)

var serverSet = wire.NewSet(wire.Value(dbType), confighelper.GetConfig, connection.OpenDBConnection, docs.DocsProvider, middleware.JwtMiddlewareProvider, middleware.DBMiddlewareProvider, validatorhelper.GetValidator, repoSet,
	serviceSet, app.CreateFiber,
)

// web dependencies
var webControllerSet = wire.NewSet(controller.WebControllerProvider, controller.AuthControllerProvider, controller.CategoryControllerProvider)

// cms dependencies
var cmsControllerSet = wire.NewSet(controller2.CmsControllerProvider, controller2.AuthControllerProvider, controller2.CategoryControllerProvider)
