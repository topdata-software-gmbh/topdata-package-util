package model

type GitRepoConfig struct {
	Name        string
	Description string
	URL         string
	PathSshKey  *string
	Branches    []string
	//	ReleaseBranches []string
}
