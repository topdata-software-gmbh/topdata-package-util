package pkg_commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
	"log"
)

var testGitCommand = &cobra.Command{
	Use:   "test-git",
	Short: "Testing git cli wrapper",
	Run: func(cmd *cobra.Command, args []string) {
		pathWebserverConfigFile, _ := cmd.Flags().GetString("WebserverConfigFile")

		fmt.Printf("Reading webserver config file: %s\n", pathWebserverConfigFile)
		webserverConfig, err := model.LoadWebserverConfig(pathWebserverConfigFile)
		if err != nil {
			log.Fatalf("Failed to load webserverConfig: %s", err)
		}

		color.Cyan("Loaded %d repository configs\n", len(webserverConfig.RepositoryConfigs))
		// iterate over the repository configs
		for _, repoConfig := range webserverConfig.RepositoryConfigs {
			color.Cyan("Cloning repository %s from %s\n", repoConfig.Name, repoConfig.URL)
			// CloneRepo the repository
			err := git_cli_wrapper.CloneRepo(repoConfig)
			if err != nil {
				log.Println("Failed to clone repository %s: [%s]: %s", repoConfig.Name, repoConfig.URL, err)
			}
			//// Fetch the branches of the repository

			color.Cyan("Fetching branches for repository %s\n", repoConfig.Name)
			branches, err := git_cli_wrapper.FetchRepoBranches(repoConfig.URL)
			if err != nil {
				log.Fatalf("Failed to fetch branches for repository %s: %s", repoConfig.Name, err)
			}
			fmt.Printf("Fetched branches for repository %s: %v\n", repoConfig.Name, branches)
		}

	},
}

func init() {
	pkgRootCommand.AddCommand(testGitCommand)
}
