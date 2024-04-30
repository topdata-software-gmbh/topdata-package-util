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

var config model.Config

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

	fmt.Printf("Loaded repositories: %+v\n", config.Repositories)
	fmt.Println("Starting server at http://localhost:" + port)
	//fmt.Println("API Endpoints:")
	//fmt.Printf("http://localhost:%s/\n", port)
	//fmt.Printf("http://localhost:%s/repositories\n", port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getRepositories(c *gin.Context) {
	repositories := git_repository_service.GetRepositories(config)
	// ---- fetch branches from the repository
	for i, repository := range repositories {
		branches, err := git_repository_service.GetRepositoryBranches(repository)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Failed to fetch branches for repository %s: %s", repository.Name, err),
			})
			return
		}
		repositories[i].Branches = branches
	}
	c.JSON(http.StatusOK, repositories)
}
