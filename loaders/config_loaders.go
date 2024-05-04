package loaders

import (
	"fmt"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/yosuke-furukawa/json5/encoding/json5"
	"io"
	"log"
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

func LoadPackagePortfolioFile(pathConfigFile string) model.PkgConfigList {
	fmt.Printf("Reading packages portfolio file: %s\n", pathConfigFile)
	var configs []model.PkgConfig
	err := loadJSONFile(pathConfigFile, &configs)
	if err != nil {
		log.Fatalln("Failed to load package portfolio", err)
	}

	return model.PkgConfigList{PkgConfigs: configs}
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
