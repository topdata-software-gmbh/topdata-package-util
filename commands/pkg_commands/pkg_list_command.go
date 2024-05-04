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
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")

		fmt.Printf("---- Reading webserver config file: %s\n", pathPackagesPortfolioFile)
		pkgConfigs, _ := loaders.LoadPackagePortfolioFile(pathPackagesPortfolioFile)

		pkgInfos := make([]model.PkgInfo, len(pkgConfigs))
		cli_out.DumpPkgsTable(pkgInfos)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgListCommand)
}
