package _struct

import (
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"io/ioutil"
	"os"
)

type Config struct {
	Repositories []GitRepository `json:"repositories"`
}

func LoadConfig(path string) (Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json5.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
