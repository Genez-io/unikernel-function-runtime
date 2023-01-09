package api

import (
	"encoding/json"
	"fmt"
	"manager/src/manage"
	"manager/src/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func RunImage(ctx echo.Context) error {
	fmt.Println("[Manager] Getting image")
	id := ctx.Param("id")
	var req models.RunImageRequest

	err := json.NewDecoder(ctx.Request().Body).Decode(&req)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Invalid request body")
	}

	// Validate step
	if err := validator.New().Struct(req); err != nil {
		fmt.Println(err)
		var rpc_err models.RunImageResponseError
		rpc_err.ID = req.ID
		rpc_err.Jsonrpc = req.Jsonrpc
		rpc_err.Error.Code = -32600
		rpc_err.Error.Message = "Invalid Request"
		return ctx.JSON(400, rpc_err)
	}

	// Search for image file
	var image_id string
	image_id, err = manage.LookupImage(req.Method, id)
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(404, err)
	}
	fmt.Println(image_id)
	res, err := manage.NewInstance(req, "test")
	if err != nil {
		fmt.Println(err)
		return ctx.JSON(404, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
