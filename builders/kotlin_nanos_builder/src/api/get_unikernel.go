package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func GetUnikernel(ctx echo.Context) error {
	fmt.Println("[Java Builder][NanOS] Getting Image")
	uuid_lookup := ctx.Param("uuid")
	image_path := "/images/" + uuid_lookup
	fmt.Println(image_path)
	return ctx.File(image_path)
}
