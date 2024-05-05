package pkg_commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
)

var buildReleaseZipCommand = &cobra.Command{
	Use:   "build-release-zip [packageName] [releaseBranchName]",
	Short: "Builds a release zip for uploading to the shopware6 plugin store",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		pkgConfig := pkgConfigList.FindOneByNameOrFail(args[0])

		packageName := args[0]
		releaseBranchName := args[1]
		color.Blue("TODO: Implement build-release-zip command for package %s and release branch %s", packageName, releaseBranchName)

		// -- switch to the release branch
		git_cli_wrapper.SwitchBranch(*pkgConfig, releaseBranchName)
		//  -- update local git repository
		// git_cli_wrapper.UpdateRepo(*pkgConfig)

		//  create a zip file
		//  upload the zip file to the shopware6 plugin store
	},
}

func init() {
	pkgRootCommand.AddCommand(buildReleaseZipCommand)
}
