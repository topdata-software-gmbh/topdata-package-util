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

// var shopwareVersion string

var findBranchCommand = &cobra.Command{
	Use:   "find-branch",
	Short: "Finds the branch with the highest plugin version for a given Shopware version",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ---- init flags
		pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)
		shopwareVersion, _ := cmd.Flags().GetString("shopware-version")

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

//func init() {
//	//findBranchCommand.Flags().StringVarP(&shopwareVersion, "shopware-version", "s", "", "Shopware version to find the branch for")
//	//_ = findBranchCommand.MarkFlagRequired("shopware-version")
//	// pkgListCommand.Flags().BoolVarP(&onlyInStore, "only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
//	pkgRootCommand.AddCommand(findBranchCommand)
//
//	pflag.StringP("shopware-version", "s", "", "Shopware version to find the branch for, eg 6.5.1")
//	pflag.BoolP("only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
//	pflag.Parse()
//	_ = viper.BindPFlags(pflag.CommandLine)
//}

func init() {
	findBranchCommand.Flags().StringP("shopware-version", "s", "", "Shopware version to find the branch for, eg 6.5.1")
	findBranchCommand.Flags().BoolP("only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
	findBranchCommand.Flags().BoolP("no-cache", "n", false, "Do not use existing cache, force rebuilding the cache")
	pkgRootCommand.AddCommand(findBranchCommand)
	//
	//viper.BindPFlag("shopware-version", findBranchCommand.Flags().Lookup("shopware-version"))
	//viper.BindPFlag("only-in-store", findBranchCommand.Flags().Lookup("only-in-store"))
	//viper.BindPFlag("no-cache", findBranchCommand.Flags().Lookup("no-cache"))
}
