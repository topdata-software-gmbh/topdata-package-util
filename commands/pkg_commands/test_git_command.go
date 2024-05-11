package pkg_commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	git_cli_wrapper2 "github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-service/globals"
	"log"
)

var testGitCommand = &cobra.Command{
	Use:   "test-git",
	Short: "Testing git cli wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		// pathPackagePortfolioFile, _ := cmd.Flags().GetString("PackagePortfolioFile")
		// pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)
		color.Cyan("Loaded %d repository configs\n", len(globals.PkgConfigList.PkgConfigs))

		// iterate over the repository configs
		for _, pkgConfig := range globals.PkgConfigList.PkgConfigs {
			color.Cyan("Cloning repository %s from %s\n", pkgConfig.Name, pkgConfig.URL)
			// CloneRepo the repository
			git_cli_wrapper2.CloneRepo(pkgConfig)
			//// Fetch the branches of the repository

			color.Cyan("Fetching branches for repository %s\n", pkgConfig.Name)
			branches, err := git_cli_wrapper2.FetchRepoBranches(pkgConfig.URL)
			if err != nil {
				log.Fatalf("Failed to fetch branches for repository %s: %s", pkgConfig.Name, err)
			}
			fmt.Printf("Fetched branches for repository %s: %v\n", pkgConfig.Name, branches)
		}

	},
}

func init() {
	pkgRootCommand.AddCommand(testGitCommand)
}
