package pkg_commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/loaders"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
	"log"
)

var testGitCommand = &cobra.Command{
	Use:   "test-git",
	Short: "Testing git cli wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("PackagesPortfolioFile")

		fmt.Printf("Reading packages portfolio file: %s\n", pathPackagesPortfolioFile)
		pkgConfigs, err := loaders.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		if err != nil {
			log.Fatalf("Failed to load packages portfolio file: %s", err)
		}

		color.Cyan("Loaded %d repository configs\n", len(pkgConfigs))
		// iterate over the repository configs
		for _, pkgConfig := range pkgConfigs {
			color.Cyan("Cloning repository %s from %s\n", pkgConfig.Name, pkgConfig.URL)
			// CloneRepo the repository
			err := git_cli_wrapper.CloneRepo(pkgConfig)
			if err != nil {
				log.Println("Failed to clone repository %s: [%s]: %s", pkgConfig.Name, pkgConfig.URL, err)
			}
			//// Fetch the branches of the repository

			color.Cyan("Fetching branches for repository %s\n", pkgConfig.Name)
			branches, err := git_cli_wrapper.FetchRepoBranches(pkgConfig.URL)
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
