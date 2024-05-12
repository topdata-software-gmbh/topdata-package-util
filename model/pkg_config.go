package model

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
	"path/filepath"
)

type PkgConfig struct {
	Name                        string
	Description                 string
	URL                         string // TODO: rename GitRepoUrl
	PathSshKey                  *string
	InShopware6Store            bool
	Shopware6StoreTechnicalName string
	Shopware6StorePluginId      int
}

func (repoConfig *PkgConfig) GetLocalGitRepoDir() string {
	return filepath.Join("/tmp/git-repos", repoConfig.Name)
}

func (repoConfig *PkgConfig) GetAbsolutePath(relativePath string) string {
	return filepath.Join(repoConfig.GetLocalGitRepoDir(), relativePath)
}

func (repoConfig *PkgConfig) IsLocalRepoExisting() bool {
	path := repoConfig.GetLocalGitRepoDir() + "/.git"
	bExists := util.FileExists(path)
	color.Blue("Checking if repo exists: %s : %t", path, bExists)

	return bExists
}
