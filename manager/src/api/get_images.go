package api

import (
	"log"
	"manager/src/models"
	"os"

	"github.com/labstack/echo/v4"
)

func GetImages(ctx echo.Context) error {
	files, err := os.ReadDir(os.Getenv("IMAGE_PATH"))
	if err != nil {
		log.Fatal(err)
		return ctx.String(500, "Something went wrong")
	}

	var response models.GetImagesSuccess = models.GetImagesSuccess{}
	for _, file := range files {
		if file.Name() == "README.md" {
			continue
		}

		response.Images = append(response.Images, file.Name())
	}

	return ctx.JSON(200, response)
}
