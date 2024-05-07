package pkg_commands

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/printer"
	"github.com/topdata-software-gmbh/topdata-package-service/serializers"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
)

var displayMode string

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		if displayMode != "compact" && displayMode != "full" {
			return fmt.Errorf("invalid displayMode value: %q, it should be either 'compact' or 'full'", displayMode)
		}
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)

		// using a cache_commands file to speed up the process
		pkgInfoList := &model.PkgInfoList{}

		if util.FileExists(config.PathCacheFile) {
			color.Yellow(">>>> Loading from cache_commands file %s", config.PathCacheFile)
			pkgInfoList = serializers.LoadPkgInfoList(config.PathCacheFile)
		} else {
			// build a list of PkgInfo objects

			pkgInfoList = factory.NewPkgInfoList(pkgConfigList)
			// save to disk for caching
			serializers.SavePkgInfoList(pkgInfoList, config.PathCacheFile)
		}

		printer.DumpPkgInfoListTable(pkgInfoList, displayMode)
		return nil
	},
}

func init() {
	pkgListCommand.Flags().StringVarP(&displayMode, "displayMode", "d", "compact", "display mode for the list, either 'compact' or 'full'")
	pkgRootCommand.AddCommand(pkgListCommand)
}
