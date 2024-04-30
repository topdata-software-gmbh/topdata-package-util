package model

type GitRepository struct {
	Name       string
	URL        string
	PathSshKey *string `json:"pathSshKey"`
	Branches   []string
	//	ReleaseBranches []string
}
