package config

import (
	"fmt"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// TODO: remove this function
//func LoadWebserverConfig(pathWebserverConfigFile string) (model.WebserverConfig, error) {
//	fmt.Printf(">>>> Reading webserver config file: %s\n", pathWebserverConfigFile)
//	var config model.WebserverConfig
//
//	viper.SetConfigFile(pathWebserverConfigFile)
//	err := viper.ReadInConfig()
//	if err != nil {
//		return config, fmt.Errorf("failed to load webserver config: %w", err)
//	}
//
//	err = viper.Unmarshal(&config)
//	if err != nil {
//		return config, fmt.Errorf("failed to unmarshal webserver config: %w", err)
//	}
//
//	return config, nil
//}

// LoadPackagePortfolioFile reads a YAML file with package definitions
func LoadPackagePortfolioFile(pathConfigFile string) *model.PkgConfigList {
	fmt.Printf(">>>> Reading portfolio file %s ... \n", pathConfigFile)

	// ---- Read the file
	data, err := ioutil.ReadFile(pathConfigFile)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	// ---- Unmarshal the data
	var portfolio model.PkgConfigList
	err = yaml.Unmarshal(data, &portfolio)
	if err != nil {
		log.Fatalf("error unmarshalling: %v", err)
	}

	fmt.Printf("Loaded %s with %d items\n", pathConfigFile, len(portfolio.Items))

	return &portfolio
}
