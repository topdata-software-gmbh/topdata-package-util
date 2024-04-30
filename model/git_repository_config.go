package model

type GitRepositoryConfig struct {
	Name       string
	URL        string
	PathSshKey *string
	Branches   []string
	//	ReleaseBranches []string
}
