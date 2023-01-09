package manage

import (
	"manager/src/models"
	"os"
)

var img_cache map[string]models.Image = nil

func InitCache() {
	img_cache = make(map[string]models.Image)

}

func grabImage(uuid string) (string, error) {
	// src := ""
	// switch lang {
	// case "js":
	// 	src = "http://localhost:12710/unikernel/"
	// case "kt":
	// 	src = "http://localhost:2709/unikernel/"
	// }
	// url := src + uuid
	// req, err := http.NewRequest(http.MethodGet, url, nil)
	// if err != nil {
	// 	fmt.Printf("client: could not create request: %s\n", err)
	// 	os.Exit(1)
	// }

	// res, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("client: error making http request: %s\n", err)
	// 	os.Exit(1)
	// }
	// image_file, err := os.Create(os.Getenv("IMAGE_PATH") + "/" + uuid)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer image_file.Close()

	// _, err = io.Copy(image_file, res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return os.Getenv("IMAGE_PATH") + "/" + uuid, nil
}

// Look for image file locally or download it
func LookupImage(uuid string, unikernel string) (string, error) {
	if _, ok := img_cache[uuid]; !ok {
		_, err := grabImage(uuid)
		img_cache[uuid] = models.Image{
			Unikernel: "nanos",
			UUID:      uuid,
		}
		if err != nil {
			return "", err
		}
	}

	return os.Getenv("IMAGE_PATH") + "/" + "kt_osv_uuid", nil
}
