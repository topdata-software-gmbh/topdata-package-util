package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/printer"
)

var displayMode string
var noCache bool

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		printer.DumpPkgInfoListTable(pkgInfoList, displayMode)
		return nil
	},
}

func init() {
	pkgListCommand.Flags().BoolVarP(&noCache, "no-cache", "n", false, "Do not use existing cache, force rebuilding the cache")
	pkgListCommand.Flags().StringVarP(&displayMode, "displayMode", "d", "compact", "display mode for the list, either 'compact' or 'full'")
	pkgRootCommand.AddCommand(pkgListCommand)
}
