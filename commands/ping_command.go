package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var pingCommand = &cobra.Command{
	Use:   "ping",
	Short: "Just a test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pong")
	},
}

func init() {
	appRootCmd.AddCommand(pingCommand)
}
