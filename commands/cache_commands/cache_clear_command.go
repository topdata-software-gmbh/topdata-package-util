package cache_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"os"
)

var cacheClearCommand = &cobra.Command{
	Use:   "clear",
	Short: "Clears the cache",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Clearing the cache %s ...\n", config.PathCacheFile)
		// delete cache file
		err := os.Remove(config.PathCacheFile)
		if err != nil {
			fmt.Printf("Failed to clear the cache: %v\n", err)
		} else {
			fmt.Printf("Cache cleared successfully\n")
		}
	},
}

func init() {
	cacheRootCommand.AddCommand(cacheClearCommand)
}
