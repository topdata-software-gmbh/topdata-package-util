package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-util/factory"
	"github.com/topdata-software-gmbh/topdata-package-util/globals"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"github.com/topdata-software-gmbh/topdata-package-util/printer"
)

//var displayMode string
//var noCache bool
// var onlyInStore bool

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		// ---- init flags
		displayMode, _ := cmd.Flags().GetString("display-mode")
		noCache, _ := cmd.Flags().GetBool("no-cache")
		onlyInStore, _ := cmd.Flags().GetBool("only-in-store")
		// 		shopwareVersion, _ := cmd.Flags().GetString("shopware-version")

		// ----
		if displayMode != "compact" && displayMode != "full" {
			return fmt.Errorf("invalid displayMode value: %q, it should be either 'compact' or 'full'", displayMode)
		}
		// pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		// pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)

		var pkgInfoList *model.PkgInfoList
		pkgInfoList = factory.NewPkgInfoListCached(globals.PkgConfigList, noCache)

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
	pkgListCommand.Flags().StringP("shopware-version", "s", "", "Shopware version to find the branch for, eg 6.5.1")
	pkgListCommand.Flags().BoolP("only-in-store", "o", false, "Show only packages that are in the Shopware6 store")
	pkgListCommand.Flags().BoolP("no-cache", "n", false, "Do not use existing cache, force rebuilding the cache")
	pkgListCommand.Flags().StringP("display-mode", "d", "compact", "display mode for the list, either 'compact' or 'full")
	pkgRootCommand.AddCommand(pkgListCommand)

	//viper.BindPFlag("shopware-version", pkgListCommand.Flags().Lookup("shopware-version"))
	//viper.BindPFlag("only-in-store", pkgListCommand.Flags().Lookup("only-in-store"))
	//viper.BindPFlag("no-cache", pkgListCommand.Flags().Lookup("no-cache"))
}
