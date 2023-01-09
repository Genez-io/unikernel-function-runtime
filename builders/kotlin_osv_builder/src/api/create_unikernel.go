package api

import (
	"encoding/json"
	"fmt"
	"kotlin_osv_builder/src/builder"
	"kotlin_osv_builder/src/models"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateUnikernel(ctx echo.Context) error {
	fmt.Println("[Java Builder][NanOS] Creating new image")
	var req models.CreateUnikernelRequest

	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Invalid request body")
	}

	// Validate step
	if err := validator.New().Struct(req); err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Invalid request body, missing fields")
	}

	res, err := builder.BuildOSvImage(req)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, res)
}
