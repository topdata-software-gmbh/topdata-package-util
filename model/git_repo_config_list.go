package model

type GitRepoConfigList struct {
	RepoConfigs []GitRepoConfig
}

func (rcl *GitRepoConfigList) FindByName(name string) *GitRepoConfig {
	for _, repoConfig := range rcl.RepoConfigs {
		if repoConfig.Name == name {
			return &repoConfig
		}
	}
	return nil
}
