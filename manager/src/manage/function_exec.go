package manage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"manager/src/models"
	"manager/src/networking"
	"net/http"
)

func ExecuteFunction(raw_req models.RunImageRequest, intf networking.TapInterface) string {

	return "Not implemented"
	json_data, err := json.Marshal(raw_req)

	if err != nil {
		log.Fatal(err)
	}

	url := "http://172.16.1.2:8000"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "curl/7.29.0")
	req.Header.Set("Accept", "*/*")

	cli := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
	return "DA"
}
