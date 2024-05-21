package git_commands

import (
	"github.com/spf13/cobra"
)

var gitRootCommand = &cobra.Command{
	Use:   "git",
	Short: "git utils for a git repo at cwd",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(gitRootCommand)
}
