package model

type GitRepositoryConfig struct {
	Name        string
	Description string
	URL         string
	PathSshKey  *string
	Branches    []string
	//	ReleaseBranches []string
}
