package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_repository_service"
	"log"
	"net/http"
)

var config model.ServiceConfig

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config.json5", "path to the config file")
}

func main() {
	flag.Parse()

	var err error
	configFile := configFile
	fmt.Printf("Reading config file: %s\n", configFile)
	config, err = model.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the TopData Package Service!")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/repositories", getRepositories)

	fmt.Printf("Loaded repositories: %+v\n", config.RepositoryConfigs)
	fmt.Println("Starting server at http://localhost:" + fmt.Sprint(config.Port))
	err = router.Run(":" + fmt.Sprint(config.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getRepositories(c *gin.Context) {
	repoInfos, err := git_repository_service.GetRepoInfos(config.RepositoryConfigs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, repoInfos)
}
