package cache_commands

import (
	"github.com/spf13/cobra"
)

var gitConfigPath string

var cacheRootCommand = &cobra.Command{
	Use:   "cache",
	Short: "cache management",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cacheRootCommand)
}
