package pkg_commands

import (
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
)

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		pkgInfos := make([]model.PkgInfo, len(pkgConfigList.PkgConfigs))

		for i, pkgConfig := range pkgConfigList.PkgConfigs {
			pkgInfos[i] = factory.NewPkgInfo(pkgConfig)
		}
		cli_out.DumpPkgsTable(pkgInfos)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgListCommand)
}
