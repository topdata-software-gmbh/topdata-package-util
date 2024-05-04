package loaders

import (
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"io"
	"os"
)

func LoadWebserverConfig(path string) (WebserverConfig, error) {
	var config WebserverConfig

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
