package git

import (
	"github.com/spf13/cobra"
)

var gitConfigPath string

var gitRootCmd = &cobra.Command{
	Use:   "git",
	Short: "git repository and branch management",
}

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(gitRootCmd)
}
