package loaders

import (
	"fmt"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"io"
	"os"
)

func LoadWebserverConfig(pathWebserverConfigFile string) (model.WebserverConfig, error) {
	var config model.WebserverConfig
	err := loadJSONFile(pathWebserverConfigFile, &config)
	if err != nil {
		return config, fmt.Errorf("failed to load webserver config: %w", err)
	}
	return config, nil
}

func LoadPackagePortfolioFile(pathConfigFile string) ([]model.PkgConfig, error) {
	var configs []model.PkgConfig
	err := loadJSONFile(pathConfigFile, &configs)
	if err != nil {
		return configs, fmt.Errorf("failed to load package portfolio: %w", err)
	}
	return configs, nil
}

func loadJSONFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json5.Unmarshal(bytes, v)
	if err != nil {
		return err
	}

	return nil
}
