package models

type CreateUnikernelRequest struct {
	Code string `json:"code"`
}

type GetUnikernelRequest struct {
	UUID string `json:"uuid"`
}