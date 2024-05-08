package pkg_commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/branches"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/printer"
)

var shopwareVersion string

var findBranchCommand = &cobra.Command{
	Use:   "find-branch",
	Short: "Finds the branch with the highest plugin version for a given Shopware version",
	RunE: func(cmd *cobra.Command, args []string) error {
		pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)

		// ---- TODO: Filter the packages based on the provided Shopware version using semantical versioning
		pkgInfoList := factory.NewPkgInfoListCached(pkgConfigList)
		ret := make([]model.PkgAndBranch, 0)
		for _, pkg := range pkgInfoList.PkgInfos {
			branch, err := branches.FindBranchForShopwareVersion(pkg, shopwareVersion)
			if err != nil {
				color.Red("FindBranchForShopwareVersion returned an error: %s\n", err.Error())
			}
			ret = append(ret, model.PkgAndBranch{Pkg: &pkg, Branch: branch})
		}

		printer.DumpPkgAndBranchTable(ret)

		return nil
	},
}

func init() {
	findBranchCommand.Flags().StringVarP(&shopwareVersion, "shopware-version", "s", "", "Shopware version to find the branch for")
	_ = findBranchCommand.MarkFlagRequired("shopware-version")

	pkgRootCommand.AddCommand(findBranchCommand)
}
