package models

type CreateUnikernelRequest struct {
	Code string `json:"code"`
	UUID string `json:"uuid"`
}

type GetUnikernelRequest struct {
	UUID string `json:"uuid"`
}