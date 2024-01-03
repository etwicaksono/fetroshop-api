// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeWebServer() error {
	viper := helper.GetConfig()
	docsDocs := docs.NewDocs(viper)
	validate := helper.GetValidator()
	userRepository := postgres.NewUserRepository()
	registrationService := registration.NewRegistrationService(userRepository)
	registrationController := controller.NewRegistrationController(validate, registrationService)
	routerRouter := router.WebRouterProvider(docsDocs, registrationController)
	serverConfig := web.WebServerConfigProvider(routerRouter)
	fiberApp := app.CreateFiber(serverConfig)
	error2 := app.StartFiber(fiberApp, serverConfig)
	return error2
}

func InitializeCmsServer() error {
	viper := helper.GetConfig()
	docsDocs := docs.NewDocs(viper)
	routerRouter := router.CmsRouterProvider(docsDocs)
	serverConfig := cms.CmsServerConfigProvider(routerRouter)
	fiberApp := app.CreateFiber(serverConfig)
	error2 := app.StartFiber(fiberApp, serverConfig)
	return error2
}

// injector.go:

var userSet = wire.NewSet(postgres.NewUserRepository, registration.NewRegistrationService, controller.NewRegistrationController)
