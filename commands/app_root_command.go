package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/commands/pkg_commands"
	"os"
)

var appRootCmd = &cobra.Command{
	Use:   "main",
	Short: "The entrypoint",
}

func Execute() {
	if err := appRootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var WebserverConfigFile string
var PackagesPortfolioFile string

func init() {
	appRootCmd.PersistentFlags().StringVar(&WebserverConfigFile, "webserver-config-file", "webserver-config.json5", "config file (default is webserver-webserver-config.json5)") // TODO: move to webserver_command.go
	appRootCmd.PersistentFlags().StringVar(&PackagesPortfolioFile, "packages-portfolio-file", "packages-portfolio.json5", "config file (default is packages-portfolio.json5)")
	pkg_commands.Register(appRootCmd)
}
