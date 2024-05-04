package pkg_commands

import (
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
)

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("PackagesPortfolioFile")
		pkgInfos := make([]model.PkgInfo, len(pathPackagesPortfolioFile))
		cli_out.DumpPkgsTable(pkgInfos)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgListCommand)
}
