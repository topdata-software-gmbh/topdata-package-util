package git_cli_wrapper

import (
	"github.com/topdata-software-gmbh/topdata-package-service/model"
)

func CloneRepo(repoConfig model.PkgConfig) {
	// Execute the pkg command to clone the repository
	folderName := repoConfig.GetLocalGitRepoDir()

	_ = execCommand("git", "clone", repoConfig.URL, folderName)
}
