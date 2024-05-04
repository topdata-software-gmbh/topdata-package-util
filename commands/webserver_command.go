package commands

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/controllers"
	"github.com/topdata-software-gmbh/topdata-package-service/gin_middleware"
	"github.com/topdata-software-gmbh/topdata-package-service/loaders"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"net/http"
)

var webserverConfig model.WebserverConfig

var (
	portFromCliOption string
)

func init() {
	flag.StringVar(&portFromCliOption, "port", "", "port to run the server on")
}

var webserverCommand = &cobra.Command{
	Use:   "webserver",
	Short: "Start the webserver",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		router := gin.Default()

		var err error

		// ---- webserver config
		fmt.Printf("---- Reading webserver config file: %s\n", WebserverConfigFile)
		webserverConfig, err = loaders.LoadWebserverConfig(WebserverConfigFile)
		if err != nil {
			log.Fatalf("Failed to load webserverConfig: %s", err)
		}

		if webserverConfig.Username != nil && webserverConfig.Password != nil {
			router.Use(gin.BasicAuth(gin.Accounts{
				*webserverConfig.Username: *webserverConfig.Password,
			}))
		}

		// pkg configs / pkg portfolio
		fmt.Printf("Reading packages portfolio file: %s\n", PackagesPortfolioFile)
		pkgConfigs, err := loaders.LoadPackagePortfolioFile(PackagesPortfolioFile)
		if err != nil {
			log.Fatalf("Failed to load packages portfolio file: %s", err)
		}

		// ---- register loaded configs in middlewares
		router.Use(gin_middleware.WebserverConfigMiddleware(webserverConfig))
		router.Use(gin_middleware.PkgConfigsMiddleware(pkgConfigs))

		// ---- define routes
		router.GET("/", welcomeHandler)
		router.GET("/ping", pingHandler)
		router.GET("/repositories", controllers.GetRepositoriesHandler)
		router.GET("/repository-details/:name", controllers.GetRepositoryDetailsHandler)

		color.Cyan("Loaded %d repository configs\n", len(pkgConfigs))

		// ---- get port
		finalPort := portFromCliOption
		if finalPort == "" {
			if webserverConfig.Port != 0 {
				finalPort = fmt.Sprint(webserverConfig.Port)
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
	appRootCmd.AddCommand(webserverCommand)
}
