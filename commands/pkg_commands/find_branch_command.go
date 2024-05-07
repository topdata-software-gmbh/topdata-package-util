package pkg_commands

import (
	"github.com/Masterminds/semver/v3"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
)

var shopwareVersion string

var findBranchCommand = &cobra.Command{
	Use:   "find-branch",
	Short: "Finds the branch with the highest plugin version for a given Shopware version",
	RunE: func(cmd *cobra.Command, args []string) error {
		pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)

		// ---- TODO: Filter the packages based on the provided Shopware version using semantical versioning
		filteredPkgInfos := make([]model.PkgInfo, 0)
		pkgInfoList := factory.NewPkgInfoListCached(pkgConfigList)
		for _, pkgInfo := range pkgInfoList.PkgInfos {
			if doesVersionConstraintMatch(pkgInfo, shopwareVersion) {
				filteredPkgInfos = append(filteredPkgInfos, pkgInfo)
			}
		}
		return nil
	},
}

func doesVersionConstraintMatch(pkgInfo model.PkgInfo, shopwareVersion string) bool {
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

	color.Cyan("Checking if %s meets %s", shopwareVersion, constraint)
	// return true // fixme

	c, err := semver.NewConstraint(constraint) // ">= 1.2.3"
	if err != nil {
		log.Fatalln("Constraint not being parsable: " + err.Error())
	}

	v, err := semver.NewVersion(shopwareVersion) // "1.1.2"
	if err != nil {
		log.Fatalln("Shopware version not being parsable: " + err.Error())
	}
	// Check if the shopwareVersion meets the constraints. The variable a will be true.
	return c.Check(v)
}

func init() {
	findBranchCommand.Flags().StringVarP(&shopwareVersion, "shopware-version", "s", "", "Shopware version to find the branch for")
	pkgRootCommand.AddCommand(findBranchCommand)
}
