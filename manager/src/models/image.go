package models

import "github.com/firecracker-microvm/firecracker-go-sdk"

type Image struct {
	Unikernel string
	UUID      string
	Config    firecracker.Config
}

type ExecutionResult struct {
	Output   string `json:"output"`
	Duration string `json:"duration"`
}
