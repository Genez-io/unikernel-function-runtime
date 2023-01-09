package main

import (
	"kotlin_nanos_builder/src/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/unikernel", api.CreateUnikernel)
	e.GET("/unikernel/:uuid", api.GetUnikernel)

	// Start server
	e.Logger.Fatal(e.Start(":2709"))

	select {}
}
