package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/topdata-software-gmbh/topdata-package-util/commands/cache_commands"
	"github.com/topdata-software-gmbh/topdata-package-util/commands/localgit_commands"
	"github.com/topdata-software-gmbh/topdata-package-util/commands/pkg_commands"
	"github.com/topdata-software-gmbh/topdata-package-util/config"
	"github.com/topdata-software-gmbh/topdata-package-util/globals"
	"os"
)

const (
	DefaultPort     = "8080"
	DefaultUsername = "default_username"
	DefaultPassword = "default_password"
)

var appRootCommand = &cobra.Command{
	Use:   "main",
	Short: "The entrypoint",
}

func Execute() {
	if err := appRootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// ---- Set up Viper
	viper.SetConfigName("topdata-package-util") // name of config file (without extension)
	viper.SetConfigType("yaml")                 // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc")          // second path to look for the config file in
	viper.AddConfigPath("$HOME/.config") // first path to look for the config file in
	// TODO... fix these hardcoded paths?
	viper.AddConfigPath("/topdata/topdata-package-util")
	viper.AddConfigPath("/topdata/topdata-package-portfolio")

	// ---- Set default values
	viper.SetDefault("webserver.port", DefaultPort)
	viper.SetDefault("webserver.username", DefaultUsername)
	viper.SetDefault("webserver.password", DefaultPassword)
	viper.SetDefault("portfolio", "/topdata/topdata-package-portfolio/portfolio.yaml")
	// Add more default values as needed

	// ---- Read from environment variables
	viper.AutomaticEnv() // automatically look for environment variables

	// ---- Read config file
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("No config file found. Using default values: %v\n", err)
	}

	// ---- Register commands
	pkg_commands.Register(appRootCommand)
	cache_commands.Register(appRootCommand)
	localgit_commands.Register(appRootCommand)

	// ---- Load package portfolio file
	globals.PkgConfigList = config.LoadPackagePortfolioFile(viper.GetString("portfolio"))
}
