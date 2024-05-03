package git_cli_wrapper

import (
	"github.com/topdata-software-gmbh/topdata-package-service/model"
)

func CloneRepo(repoConfig model.GitRepoConfig) error {
	// Execute the git command to clone the repository
	folderName := GetLocalGitRepoDir(repoConfig)

	err2 := execCommand("git", "clone", repoConfig.URL, folderName)
	if err2 != nil {
		return err2
	}

	return nil
}
