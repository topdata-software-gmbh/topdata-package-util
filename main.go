package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/pkg"
	"log"
	"net/http"
)

var config pkg.Config

var (
	port       string
	configFile string
)

func init() {
	flag.StringVar(&port, "port", "8080", "port to run the server on")
	flag.StringVar(&configFile, "config", "config.json5", "path to the config file")
}

func main() {
	flag.Parse()

	var err error
	configFile := configFile
	fmt.Printf("Reading config file: %s\n", configFile)
	config, err = pkg.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the TopData Package Service!")
	})

	router.GET("/repositories", getRepositories)

	fmt.Printf("Loaded repositories: %+v\n", config.Repositories)
	fmt.Printf("Server started at http://localhost:%s\n", port)
	fmt.Println("API Endpoints:")
	fmt.Printf("http://localhost:%s/\n", port)
	fmt.Printf("http://localhost:%s/repositories\n", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getRepositories(c *gin.Context) {
	c.JSON(http.StatusOK, config.Repositories)
}
