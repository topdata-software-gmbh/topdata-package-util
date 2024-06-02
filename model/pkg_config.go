package model

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
	"path/filepath"
)

type PkgConfig struct {
	Name                        string `yaml:"name"`
	Description                 string `yaml:"description"`
	URL                         string `yaml:"url"`
	PathSshKey                  string `yaml:"pathSshKey"`
	InShopware6Store            bool   `yaml:"inShopware6Store"`
	Shopware6StorePluginId      int    `yaml:"shopware6StorePluginId"`
	Shopware6StoreTechnicalName string `yaml:"shopware6StoreTechnicalName"`
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
