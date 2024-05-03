package git_service_v2

import (
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/file_path_service"
)

func CloneRepo(repoConfig model.GitRepoConfig) error {
	// Execute the git command to clone the repository
	folderName := file_path_service.GetLocalGitRepoDir(repoConfig)

	err2 := execCommand("git", "clone", repoConfig.URL, folderName)
	if err2 != nil {
		return err2
	}

	return nil
}
