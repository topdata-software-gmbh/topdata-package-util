package pkg_commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	git_cli_wrapper2 "github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
	"log"
)

var testGitCommand = &cobra.Command{
	Use:   "test-git",
	Short: "Testing git cli wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("PackagesPortfolioFile")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		color.Cyan("Loaded %d repository configs\n", len(pkgConfigList.PkgConfigs))

		// iterate over the repository configs
		for _, pkgConfig := range pkgConfigList.PkgConfigs {
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
