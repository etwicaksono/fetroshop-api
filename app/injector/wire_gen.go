// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	postgres2 "github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses/postgres"
	postgres3 "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses/postgres"
	postgres4 "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	controller2 "github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	auth2 "github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeWebServer() *app.Fetroshop {
	viper := confighelper.GetConfig()
	docsDocs := docs.DocsProvider(viper)
	connectionDBType := _wireDBTypeValue
	connectionConnection := connection.OpenDBConnection(connectionDBType, viper)
	userAccessRepo := postgres.RepoProvider(connectionConnection)
	customerAccessRepo := postgres2.RepoProvider(connectionConnection)
	jwtMiddleware := middleware.JwtMiddlewareProvider(viper, userAccessRepo, customerAccessRepo)
	dbMiddleware := middleware.DBMiddlewareProvider(connectionConnection)
	validate := validatorhelper.GetValidator()
	customerRepo := postgres3.RepoProvider(connectionConnection)
	authService := auth.ServiceProvider(connectionConnection, viper, validate, customerRepo, customerAccessRepo)
	authController := controller.AuthControllerProvider(validate, authService)
	controllerController := controller.WebControllerProvider(authController)
	routerRouter := router.WebRouterProvider(docsDocs, jwtMiddleware, dbMiddleware, controllerController)
	serverConfig := web.WebServerConfigProvider(routerRouter)
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
	connectionConnection := connection.OpenDBConnection(connectionDBType, viper)
	userAccessRepo := postgres.RepoProvider(connectionConnection)
	customerAccessRepo := postgres2.RepoProvider(connectionConnection)
	jwtMiddleware := middleware.JwtMiddlewareProvider(viper, userAccessRepo, customerAccessRepo)
	dbMiddleware := middleware.DBMiddlewareProvider(connectionConnection)
	validate := validatorhelper.GetValidator()
	userRepo := postgres4.RepoProvider(connectionConnection)
	authService := auth2.ServiceProvider(connectionConnection, viper, validate, userRepo, userAccessRepo)
	authController := controller2.AuthControllerProvider(authService)
	controllerController := controller2.CmsControllerProvider(authController)
	routerRouter := router.CmsRouterProvider(docsDocs, jwtMiddleware, dbMiddleware, controllerController)
	serverConfig := cms.CmsServerConfigProvider(routerRouter)
	fetroshop := app.CreateFiber(serverConfig)
	return fetroshop
}

// injector.go:

var dbType connection.DBType = connection.DB_MAIN

var serverSet = wire.NewSet(wire.Value(dbType), confighelper.GetConfig, connection.OpenDBConnection, docs.DocsProvider, middleware.JwtMiddlewareProvider, middleware.DBMiddlewareProvider, validatorhelper.GetValidator, postgres.RepoProvider, postgres2.RepoProvider, app.CreateFiber)

// web dependencies
var webRepoSet = wire.NewSet(postgres3.RepoProvider)

var webControllerSet = wire.NewSet(controller.WebControllerProvider, controller.AuthControllerProvider)

var webServiceSet = wire.NewSet(auth.ServiceProvider)

// cms dependencies
var cmsRepoSet = wire.NewSet(postgres4.RepoProvider)

var cmsControllerSet = wire.NewSet(controller2.CmsControllerProvider, controller2.AuthControllerProvider)

var cmsServiceSet = wire.NewSet(auth2.ServiceProvider)
