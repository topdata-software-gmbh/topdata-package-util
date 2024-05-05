package model

import (
	"os"
	"path/filepath"
)

type PkgConfig struct {
	Name        string
	Description string
	URL         string // TODO: rename GitRepoUrl
	PathSshKey  *string
	Branches    []string
	//	ReleaseBranches []string
}

func (repoConfig *PkgConfig) GetLocalGitRepoDir() string {
	return filepath.Join("/tmp/git-repos", repoConfig.Name)
}

func (repoConfig *PkgConfig) GetAbsolutePath(relativePath string) string {
	return filepath.Join(repoConfig.GetLocalGitRepoDir(), relativePath)
}

func (repoConfig *PkgConfig) IsLocalRepoExisting() bool {
	path := repoConfig.GetLocalGitRepoDir() + "/.git"
	_, err := os.Stat(path)

	return os.IsExist(err)
}
