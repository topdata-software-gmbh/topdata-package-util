package pkg_commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var buildReleaseZipCommand = &cobra.Command{
	Use:   "build-release-zip [packageName] [releaseBranchName]",
	Short: "Builds a release zip for uploading to the shopware6 plugin store",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		releaseBranchName := args[1]
		color.Blue("TODO: Implement build-release-zip command for package %s and release branch %s", packageName, releaseBranchName)
	},
}

func init() {
	pkgRootCommand.AddCommand(buildReleaseZipCommand)
}
