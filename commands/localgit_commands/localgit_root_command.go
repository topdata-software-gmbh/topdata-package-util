package localgit_commands

import (
	"github.com/spf13/cobra"
)

/*
 * "localgit" commands are used as a helper when working with a git repo at cwd
 */

var localgitRootCommand = &cobra.Command{
	Use:   "localgit",
	Short: "git utils for a git repo at cwd",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(localgitRootCommand)
}
