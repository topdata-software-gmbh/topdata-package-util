package model

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
	"path/filepath"
)

type PkgConfig struct {
	Name                        string
	Description                 string
	URL                         string // TODO: rename GitRepoURL
	PathSshKey                  *string
	InShopware6Store            bool
	Shopware6StoreTechnicalName string
	Shopware6StorePluginId      int
}

func (pkgConfig *PkgConfig) GetLocalGitRepoDir() string {
	return filepath.Join("/tmp/git-repos", pkgConfig.Name)
}

func (pkgConfig *PkgConfig) GetAbsolutePath(relativePath string) string {
	return filepath.Join(pkgConfig.GetLocalGitRepoDir(), relativePath)
}

func (pkgConfig *PkgConfig) IsLocalRepoExisting() bool {
	path := pkgConfig.GetLocalGitRepoDir() + "/.git"
	bExists := util.FileExists(path)
	color.Blue("Checking if repo exists: %s : %t", path, bExists)

	return bExists
}
