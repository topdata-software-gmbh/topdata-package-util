package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/commands/pkg_commands"
	"os"
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

var PackagesPortfolioFile string

func init() {
	appRootCommand.PersistentFlags().StringVar(&PackagesPortfolioFile, "packages-portfolio-file", "packages-portfolio.json5", "config file (default is packages-portfolio.json5)")
	pkg_commands.Register(appRootCommand)
}
