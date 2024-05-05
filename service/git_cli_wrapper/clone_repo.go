package git_cli_wrapper

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
)

func CloneRepo(repoConfig model.PkgConfig) {
	// Execute the git command to clone the repository
	folderName := repoConfig.GetLocalGitRepoDir()

	_ = execCommand("git", "clone", repoConfig.URL, folderName)
}

func DownsyncRepo(repoConfig model.PkgConfig) {
	// check if repo exists
	if !repoConfig.IsLocalRepoExisting() {
		color.Green("Cloning repo: %s", repoConfig.Name)
		CloneRepo(repoConfig)
	} else {
		_ = execGitCommand(repoConfig, "pull")
	}
}
