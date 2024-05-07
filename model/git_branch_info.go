package model

type GitBranchInfo struct {
	Name                      string
	CommitId                  string
	PackageVersion            string // from composer.json
	ShopwareVersionConstraint string // from composer.json
}
