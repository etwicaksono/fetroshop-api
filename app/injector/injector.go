//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	customerRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	userRepo "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	cmsController "github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	cmsAuthService "github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	webController "github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	webAuthService "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	confighelper.GetConfig,
	connection.OpenDBConnection,
	docs.DocsProvider,
	middleware.JwtMiddlewareProvider,
	validatorhelper.GetValidator,
	app.CreateFiber,
	app.StartFiber,
)

// web dependencies
var webRepoSet = wire.NewSet(
	customerRepo.CustomerRepositoryProvider,
)

var webControllerSet = wire.NewSet(
	webController.WebControllerProvider,
	webController.AuthControllerProvider,
)

var webServiceSet = wire.NewSet(
	webAuthService.AuthServiceProvider,
)

func InitializeWebServer() error {
	wire.Build(
		serverSet,
		webRepoSet,
		webControllerSet,
		webServiceSet,
		router.WebRouterProvider,
		web.WebServerConfigProvider,
	)
	return nil
}

// cms dependencies
var cmsRepoSet = wire.NewSet(
	userRepo.UserRepositoryProvider,
)

var cmsControllerSet = wire.NewSet(
	cmsController.CmsControllerProvider,
	cmsController.AuthControllerProvider,
)

var cmsServiceSet = wire.NewSet(
	cmsAuthService.AuthServiceProvider,
)

func InitializeCmsServer() error {
	wire.Build(
		serverSet,
		cmsRepoSet,
		cmsControllerSet,
		cmsServiceSet,
		router.CmsRouterProvider,
		cms.CmsServerConfigProvider,
	)
	return nil
}
