package main

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/injector"
)

func main() {

	// Run cms server
	go func() {
		cmsApp := injector.InitializeCmsServer()
		err := cmsApp.FiberApp.Listen(fmt.Sprintf("%s:%d", cmsApp.Host, cmsApp.Port))
		if err != nil {
			cmsApp.Logger.Panic(err.Error())
		}
	}()

	// Run web server
	webApp := injector.InitializeWebServer()
	err := webApp.FiberApp.Listen(fmt.Sprintf("%s:%d", webApp.Host, webApp.Port))
	if err != nil {
		webApp.Logger.Panic(err.Error())
	}
}
