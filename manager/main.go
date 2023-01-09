package main

import (
	"manager/src/api"
	"manager/src/manage"
	"manager/src/networking"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init Networking Package
	networking.InitNetworking()
	manage.InitCache()
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("api/images/:id", api.RunImage)
	e.GET("api/images", api.GetImages)

	// Start server
	e.Logger.Fatal(e.Start(":2806"))

	select {}
}
