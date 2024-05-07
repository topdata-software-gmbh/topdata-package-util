package config

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
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

func LoadPackagePortfolioFile(pathConfigFile string) model.PkgConfigList {
	color.Yellow(">>>> Reading packages portfolio file... \n")
	var configs []model.PkgConfig

	viper.SetConfigName("packages-portfolio") // name of config file (without extension)
	viper.SetConfigType("yaml")               // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
	viper.AddConfigPath(".") // optionally look for config in the working directory
	//err := viper.ReadInConfig() // Find and read the config file

	err := viper.Unmarshal(&configs)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	//if err != nil {             // Handle errors reading the config file
	//	panic(fmt.Errorf("Fatal reading portfolio file: %w", err))
	//}

	return model.PkgConfigList{PkgConfigs: configs}
}
