package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "The entrypoint",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config-file", "config.json5", "config file (default is config.json5)")
}
