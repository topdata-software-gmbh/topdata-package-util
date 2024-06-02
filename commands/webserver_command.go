package commands

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/topdata-software-gmbh/topdata-package-util/controllers"
	"github.com/topdata-software-gmbh/topdata-package-util/gin_middleware"
	"github.com/topdata-software-gmbh/topdata-package-util/globals"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"log"
	"net/http"
)

var (
	portFromCliOption string
)

var webserverCommand = &cobra.Command{
	Use:   "webserver",
	Short: "Start the webserver",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		router := gin.Default()

		var err error

		// ---- load webserver webserverConfig

		var webserverConfig model.WebserverConfig
		err = viper.UnmarshalKey("webserver", &webserverConfig)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}

		fmt.Printf("WebserverConfig: %+v\n", webserverConfig)
		if webserverConfig.Username != "" && webserverConfig.Password != "" {
			router.Use(gin.BasicAuth(gin.Accounts{webserverConfig.Username: webserverConfig.Password}))
		}

		// pkgConfigList := webserverConfig.LoadPackagePortfolioFile(PackagePortfolioFile)

		// ---- register loaded configs in middlewares
		router.Use(gin_middleware.WebserverConfigMiddleware(webserverConfig))
		router.Use(gin_middleware.PkgConfigListMiddleware(globals.PkgConfigList))

		// ---- define routes
		router.GET("/", welcomeHandler)
		router.GET("/ping", pingHandler)
		router.GET("/repositories", controllers.GetRepositoriesHandler)
		router.GET("/repository-details/:name", controllers.GetRepositoryDetailsHandler)

		// ----
		color.Cyan("Loaded %d repository configs\n", len(globals.PkgConfigList.Items))

		// ---- get port (TODO: remove, use spf13/viper)
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
	appRootCommand.AddCommand(webserverCommand)
}
