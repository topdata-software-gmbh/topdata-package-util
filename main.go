package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/controllers"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"net/http"
)

var serviceConfig model.ServiceConfig

var (
	portFromCliOption string
	configFile        string
)

func init() {
	flag.StringVar(&portFromCliOption, "port", "", "port to run the server on")
	flag.StringVar(&configFile, "config-file", "config.json5", "path to the service config file")
}

func ServiceConfigMiddleware(config model.ServiceConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("serviceConfig", config)
		c.Next()
	}
}

func main() {
	flag.Parse()

	var err error
	fmt.Printf("Reading serviceConfig file: %s\n", configFile)
	serviceConfig, err = model.LoadServiceConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load serviceConfig: %s", err)
	}

	router := gin.Default()
	if serviceConfig.Username != nil && serviceConfig.Password != nil {
		router.Use(gin.BasicAuth(gin.Accounts{
			*serviceConfig.Username: *serviceConfig.Password,
		}))
	}

	router.Use(ServiceConfigMiddleware(serviceConfig))

	router.GET("/", welcomeHandler)
	router.GET("/ping", pingHandler)
	router.GET("/repositories", controllers.GetRepositoriesHandler)
	router.GET("/repository-details/:name", controllers.GetRepositoryDetailsHandler)

	fmt.Printf("Loaded %d repository configs\n", len(serviceConfig.RepositoryConfigs))

	// ---- get port
	finalPort := portFromCliOption
	if finalPort == "" {
		if serviceConfig.Port != 0 {
			finalPort = fmt.Sprint(serviceConfig.Port)
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

//func getRepoNames(repoConfigs []model.GitRepositoryConfig) []string {
//	names := make([]string, len(repoConfigs))
//	for i, config := range repoConfigs {
//		names[i] = config.Name
//	}
//	return names
//}

func welcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the TopData Package Service!")
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
