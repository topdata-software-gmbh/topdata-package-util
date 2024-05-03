package model

// this is a model for the git repository info which is extracted from a git repository
// maybe also the git url? needed?
type GitRepoInfo struct {
	Name            string
	Description     string
	URL             string // optional [TODO: we want a setting whether to show this or not]
	Branches        []string
	ReleaseBranches []GitBranchInfo
}
