package commands

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/controllers"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"net/http"
)

var serviceConfig model.ServiceConfig

var (
	portFromCliOption string
)

func init() {
	flag.StringVar(&portFromCliOption, "port", "", "port to run the server on")
}

func ServiceConfigMiddleware(config model.ServiceConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("serviceConfig", config)
		c.Next()
	}
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
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

		color.Cyan("Loaded %d repository configs\n", len(serviceConfig.RepositoryConfigs))

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
	},
}

func welcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the TopData Package Service!")
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func init() {
	rootCmd.AddCommand(serverCommand)
}
