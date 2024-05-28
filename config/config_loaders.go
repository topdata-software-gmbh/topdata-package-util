package config

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"log"
)

func LoadWebserverConfig(pathWebserverConfigFile string) (model.WebserverConfig, error) {
	fmt.Printf(">>>> Reading webserver config file: %s\n", pathWebserverConfigFile)
	var config model.WebserverConfig

	viper.SetConfigFile(pathWebserverConfigFile)
	err := viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to load webserver config: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal webserver config: %w", err)
	}

	return config, nil
}

func LoadPackagePortfolioFile(pathConfigFile string) *model.PkgConfigList {
	color.Yellow(">>>> XXX Reading portfolio file %s ... \n", pathConfigFile)

	var configs []model.PkgConfig

	viper.AddConfigPath(".")
	// TODO... fix these hardcoded paths?
	viper.AddConfigPath("/topdata/topdata-package-util")
	viper.AddConfigPath("/topdata/topdata-package-portfolio")

	viper.SetConfigFile(pathConfigFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}
	if err := viper.UnmarshalKey("items", &configs); err != nil {
		log.Fatalf("error unmarshalling: %v", err)
	}
	machineName := viper.GetString("machineName")
	//	fmt.Print("Loaded " + pathConfigFile + "with " + len(configs) + " items\n")
	fmt.Printf("Loaded %s with %d items\n", pathConfigFile, len(configs))

	return &model.PkgConfigList{
		MachineName: machineName,
		PkgConfigs:  configs,
	}
}
