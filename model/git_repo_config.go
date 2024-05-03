package model

import (
	"path/filepath"
)

type GitRepoConfig struct {
	Name        string
	Description string
	URL         string
	PathSshKey  *string
	Branches    []string
	//	ReleaseBranches []string
}

func (repoConfig *GitRepoConfig) GetLocalGitRepoDir() string {
	return filepath.Join("/tmp/git-repos", repoConfig.Name)
}

func (repoConfig *GitRepoConfig) GetAbsolutePath(relativePath string) string {
	return filepath.Join(repoConfig.GetLocalGitRepoDir(), relativePath)
}
