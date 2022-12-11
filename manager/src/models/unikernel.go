package models

type UnikernelConfig struct {
	Program string            `json:"Program"`
	Args    []string          `json:"Args"`
	Files   []string          `json:"Files"`
	MapDirs map[string]string `json:"MapDirs"`
}
type Unikernel struct {
	ConfigPath string
	UUID       string
	KernelImg  string
	RootFsImg  string
	CreatedBy  string
	Config     UnikernelConfig
}
