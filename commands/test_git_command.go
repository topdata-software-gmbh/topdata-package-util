package commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_service_v2"
	"log"
)

var testGitCommand = &cobra.Command{
	Use:   "test-git",
	Short: "Testing git cli wrapper",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Reading serviceConfig file: %s\n", configFile)
		serviceConfig, err := model.LoadServiceConfig(configFile)
		if err != nil {
			log.Fatalf("Failed to load serviceConfig: %s", err)
		}

		color.Cyan("Loaded %d repository configs\n", len(serviceConfig.RepositoryConfigs))
		// iterate over the repository configs
		for _, repoConfig := range serviceConfig.RepositoryConfigs {
			color.Cyan("Cloning repository %s from %s\n", repoConfig.Name, repoConfig.URL)
			// CloneRepository the repository
			err := git_service_v2.CloneRepository(repoConfig)
			if err != nil {
				log.Println("Failed to clone repository %s: [%s]: %s", repoConfig.Name, repoConfig.URL, err)
			}
			//// Fetch the branches of the repository

			color.Cyan("Fetching branches for repository %s\n", repoConfig.Name)
			branches, err := git_service_v2.FetchRepositoryBranches(repoConfig.URL)
			if err != nil {
				log.Fatalf("Failed to fetch branches for repository %s: %s", repoConfig.Name, err)
			}
			fmt.Printf("Fetched branches for repository %s: %v\n", repoConfig.Name, branches)
		}

	},
}

func init() {
	rootCmd.AddCommand(testGitCommand)
}
