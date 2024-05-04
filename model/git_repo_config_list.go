package model

type GitRepoConfigList struct {
	RepoConfigs []PkgConfig
}

func (rcl *GitRepoConfigList) FindOneByName(name string) *PkgConfig {
	for _, repoConfig := range rcl.RepoConfigs {
		if repoConfig.Name == name {
			return &repoConfig
		}
	}
	return nil
}
