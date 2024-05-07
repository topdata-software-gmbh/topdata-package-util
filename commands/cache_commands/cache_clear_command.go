package cache_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/app_constants"
	"os"
)

var cacheClearCommand = &cobra.Command{
	Use:   "clear",
	Short: "Clears the cache",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Clearing the cache %s ...\n", app_constants.PathCacheFile)
		// delete cache file
		err := os.Remove(app_constants.PathCacheFile)
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
