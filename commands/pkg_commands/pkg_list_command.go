package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/loaders"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
)

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("PackagesPortfolioFile")

		fmt.Printf("Reading webserver config file: %s\n", pathPackagesPortfolioFile)
		pkgConfigs, err := loaders.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		if err != nil {
			fmt.Printf("Failed to load package portfolio: %s", err)
		}

		pkgInfos := make([]model.PkgInfo, len(pkgConfigs))
		cli_out.DumpPkgsTable(pkgInfos)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgListCommand)
}
