package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	appRootCommand.AddCommand(versionCommand)
}

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Topdata Package Service",
	Long:  `All software has versions. This is Topdata Package Service's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Topdata Package Service v0.1")
	},
}
