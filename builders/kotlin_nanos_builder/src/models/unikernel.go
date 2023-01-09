package models

type UnikernelConfig struct {
	Program string            `json:"Program,omitempty"`
	Args    []string          `json:"Args,omitempty"`
	Files   []string          `json:"Files,omitempty"`
	Dirs    []string          `json:"Dirs,omitempty"`
	MapDirs map[string]string `json:"MapDirs,omitempty"`
}
type Unikernel struct {
	ConfigPath string
	UUID       string
	KernelImg  string
	RootFsImg  string
	CreatedBy  string
	Config     UnikernelConfig
}
