package branches

import (
	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/factory"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"log"
)

func DoesVersionConstraintMatch(pkgInfo model.PkgInfo, shopwareVersion string) bool {
	// iterate release branches
	for _, branchName := range pkgInfo.ReleaseBranchNames {
		// get version from branch name
		gitBranchInfo := factory.NewGitBranchInfo(pkgInfo.PkgConfig, branchName)
		// check if the version matches the constraint
		if isVersionMatchingConstraint(gitBranchInfo.ShopwareVersionConstraint, shopwareVersion) {
			return true
		}
	}
	return false
}

func isVersionMatchingConstraint(constraint string, shopwareVersion string) bool {
	if constraint == "" {
		color.Red("No constraint found. Assuming it matches.")
		return true
	}

	c, err := semver.NewConstraint(constraint) // ">= 1.2.3"
	if err != nil {
		log.Fatalln("Constraint not being parsable: " + err.Error())
	}

	v, err := semver.NewVersion(shopwareVersion) // "1.1.2"
	if err != nil {
		log.Fatalln("Shopware version not being parsable: " + err.Error())
	}

	if c.Check(v) {
		color.Green(" %s meets %s", shopwareVersion, constraint)
		return true
	} else {
		color.Red(" %s does not meet %s", shopwareVersion, constraint)
		return false
	}
}
