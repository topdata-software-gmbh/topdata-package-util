package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/cli_out"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
)

var displayMode string

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	RunE: func(cmd *cobra.Command, args []string) error {
		if displayMode != "compact" && displayMode != "full" {
			return fmt.Errorf("invalid displayMode value: %q, it should be either 'compact' or 'full'", displayMode)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		pkgInfos := make([]model.PkgInfo, len(pkgConfigList.PkgConfigs))

		for i, pkgConfig := range pkgConfigList.PkgConfigs {
			pkgInfos[i] = factory.NewPkgInfo(pkgConfig)
		}
		cli_out.DumpPkgsTable(pkgInfos, displayMode)
	},
}

func init() {
	pkgListCommand.Flags().StringVarP(&displayMode, "displayMode", "d", "compact", "display mode for the list, either 'compact' or 'full'")
	pkgRootCommand.AddCommand(pkgListCommand)
}
