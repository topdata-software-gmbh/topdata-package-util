package model

type GitBranchInfo struct {
	Name                      string
	PackageVersion            string // from composer.json
	ShopwareVersionConstraint string // from composer.json
	CommitId                  string
	CommitIdShort             string
	CommitDate                string
	CommitAuthor              string
	CommitMessage             string
}
