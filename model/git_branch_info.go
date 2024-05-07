package model

type GitBranchInfo struct {
	Name                      string
	CommitId                  string
	CommitDate                string
	CommitAuthor              string
	PackageVersion            string // from composer.json
	ShopwareVersionConstraint string // from composer.json
}
