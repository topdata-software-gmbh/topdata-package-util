package model

// this is a model for the pkg repository info which is extracted from a pkg repository
// maybe also the pkg url? needed?
type PkgInfo struct {
	PkgConfig          *PkgConfig
	ReleaseBranchNames []string
	OtherBranchNames   []string
	LastCommitToMainAt string
	// Description        string
	// URL                string // optional [TODO: we want a setting whether to show this or not]
	// Name               string
	// BranchNames        []string
	// ReleaseBranches    []GitBranchInfo
}
