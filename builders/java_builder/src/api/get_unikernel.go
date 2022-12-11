package api

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func GetUnikernel(ctx echo.Context) error {
	fmt.Println("[Java Builder][NanOS] Getting Image")
	uuid_lookup := ctx.Param("uuid")
	image_path := os.Getenv("IMAGES_PATH") + uuid_lookup
	return ctx.File(image_path)
}
