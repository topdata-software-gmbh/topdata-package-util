package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config-file", "config.json5", "path to the service config file")
}

func main() {
	flag.Parse()

	fmt.Printf("Reading serviceConfig file: %s\n", configFile)
	serviceConfig, err := model.LoadServiceConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load serviceConfig: %s", err)
	}

	color.Cyan("Loaded %d repository configs\n", len(serviceConfig.RepositoryConfigs))
}
