package model

type WebserverConfig struct {
	Port     uint16 `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}
