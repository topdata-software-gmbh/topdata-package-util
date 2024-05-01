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
	portFromCliOption string
	configFile        string
)

func init() {
	flag.StringVar(&portFromCliOption, "port", "", "port to run the server on")
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
	if config.Username != nil && config.Password != nil {
		router.Use(gin.BasicAuth(gin.Accounts{
			*config.Username: *config.Password,
		}))
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the TopData Package Service!")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/repositories", getRepositoriesAction)

	fmt.Printf("Loaded %d repository configs: %v\n", len(config.RepositoryConfigs), getRepoNames(config.RepositoryConfigs))

	// ---- get port
	finalPort := portFromCliOption
	if finalPort == "" {
		if config.Port != 0 {
			finalPort = fmt.Sprint(config.Port)
		} else {
			finalPort = "8080"
		}
	}

	// ---- start the server
	fmt.Println("Starting server at http://localhost:" + finalPort)
	err = router.Run(":" + finalPort)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getRepositoriesAction(c *gin.Context) {
	repoInfos, err := git_repository_service.GetRepoInfos(config.RepositoryConfigs, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, repoInfos)
}

func getRepoNames(repoConfigs []model.GitRepositoryConfig) []string {
	names := make([]string, len(repoConfigs))
	for i, config := range repoConfigs {
		names[i] = config.Name
	}
	return names
}
