package pkg_commands

import (
	"github.com/spf13/cobra"
)

var gitConfigPath string

var pkgRootCommand = &cobra.Command{
	Use:   "pkg",
	Short: "pkg repository and branch management",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(pkgRootCommand)
}
