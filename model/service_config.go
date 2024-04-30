package model

import (
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"io"
	"os"
)

type ServiceConfig struct {
	Port              uint16                `json:"port"`
	RepositoryConfigs []GitRepositoryConfig `json:"repositories"`
}

func LoadConfig(path string) (ServiceConfig, error) {
	var config ServiceConfig

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json5.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
