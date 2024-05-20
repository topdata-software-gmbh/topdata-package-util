package branches

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/topdata-software-gmbh/topdata-package-util/factory"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"sort"
)

// FindBranchForShopwareVersion finds the branch with the highest plugin version for a given Shopware version
func FindBranchForShopwareVersion(pkgInfo model.PkgInfo, version string) (*model.GitBranchInfo, error) {
	// iterate release branches
	matchingBranches := []model.GitBranchInfo{}
	for _, branchName := range pkgInfo.ReleaseBranchNames {

		gitBranchInfo := factory.NewGitBranchInfo(pkgInfo.PkgConfig, branchName)
		// check if the version matches the constraint
		if isVersionMatchingConstraint(gitBranchInfo.ShopwareVersionConstraint, version) {
			matchingBranches = append(matchingBranches, gitBranchInfo)
		}
	}
	if len(matchingBranches) > 0 {
		// sort by PackageVersion
		sort.Slice(matchingBranches, func(i, j int) bool {
			versionI, errI := semver.NewVersion(matchingBranches[i].PackageVersion)
			versionJ, errJ := semver.NewVersion(matchingBranches[j].PackageVersion)

			if errI != nil || errJ != nil {
				// Handle version not being parsable.
				return false
			}

			// Use the CompareTo method to compare the versions
			return versionI.GreaterThan(versionJ)
		})
		return &matchingBranches[0], nil

	} else {
		return nil, fmt.Errorf("No matching branch of %s found for Shopware version %s", pkgInfo.PkgConfig.Name, version)
	}
}
