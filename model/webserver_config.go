package model

type WebserverConfig struct {
	Port     uint16  `json:"port"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	//	RepositoryConfigs []PkgConfig `json:"repositories"`
}
