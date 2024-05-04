package model

type GitBranchInfo struct {
	Name            string
	CommitId        string
	PackageVersion  string // from composer.json
	ShopwareVersion string // from composer.json

}
