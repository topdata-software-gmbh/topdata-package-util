package file_path_service

import (
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"path/filepath"
)

func GetLocalGitRepoDir(repoConf model.GitRepositoryConfig) string {
	return filepath.Join("/tmp/git-repos", repoConf.Name)
}
