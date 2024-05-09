package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/printer"
)

//var displayMode string
//var noCache bool
// var onlyInStore bool

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ---- init flags
		displayMode := viper.GetString("display-mode")
		noCache := viper.GetBool("no-cache")
		onlyInStore := viper.GetBool("only-in-store")
		// ----
		if displayMode != "compact" && displayMode != "full" {
			return fmt.Errorf("invalid displayMode value: %q, it should be either 'compact' or 'full'", displayMode)
		}
		pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)

		var pkgInfoList *model.PkgInfoList
		if noCache {
			pkgInfoList = factory.NewPkgInfoList(pkgConfigList)
		} else {
			pkgInfoList = factory.NewPkgInfoListCached(pkgConfigList)
		}

		// ---- filter
		if onlyInStore {
			pkgInfoList = pkgInfoList.FilterInShopware6Store()
		}

		printer.DumpPkgInfoListTable(pkgInfoList, displayMode)
		return nil
	},
}

/*func init() {
	//pkgListCommand.Flags().BoolVarP(&noCache, "no-cache", "n", false, "Do not use existing cache, force rebuilding the cache")
	//pkgListCommand.Flags().BoolVarP(&onlyInStore, "only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
	//pkgListCommand.Flags().StringVarP(&displayMode, "displayMode", "d", "compact", "display mode for the list, either 'compact' or 'full'")
	pkgRootCommand.AddCommand(pkgListCommand)
	pflag.BoolP("no-cache", "n", false, "Do not use existing cache, force rebuilding the cache")
	pflag.StringP("shopware-version", "s", "", "Shopware version to find the branch for, eg 6.5.1")
	pflag.BoolP("only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

}
*/

func init() {
	findBranchCommand.Flags().StringP("shopware-version", "s", "", "Shopware version to find the branch for, eg 6.5.1")
	findBranchCommand.Flags().BoolP("only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
	findBranchCommand.Flags().BoolP("no-cache", "n", false, "Do not use existing cache, force rebuilding the cache")
	pkgRootCommand.AddCommand(findBranchCommand)

	viper.BindPFlag("shopware-version", findBranchCommand.Flags().Lookup("shopware-version"))
	viper.BindPFlag("only-in-store", findBranchCommand.Flags().Lookup("only-in-store"))
	viper.BindPFlag("no-cache", findBranchCommand.Flags().Lookup("no-cache"))
}
