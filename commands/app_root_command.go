package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/commands/git"
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

var ConfigFile string

func init() {
	appRootCmd.PersistentFlags().StringVar(&ConfigFile, "config-file", "config.json5", "config file (default is config.json5)")
	git.Register(appRootCmd)
}
