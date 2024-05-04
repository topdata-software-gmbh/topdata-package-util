package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/loaders"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
)

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")

		fmt.Printf("Reading packages portfolio file: %s\n", pathPackagesPortfolioFile)
		pkgConfigs, _ := loaders.LoadPackagePortfolioFile(pathPackagesPortfolioFile)

		pkgInfos := make([]model.PkgInfo, len(pkgConfigs))
		for i, pkgConfig := range pkgConfigs {
			pkgInfos[i] = model.PkgInfo{
				Name:               pkgConfig.Name,
				URL:                pkgConfig.URL,
				BranchNames:        git_cli_wrapper.GetBranchNames(pkgConfig),
				ReleaseBranchNames: git_cli_wrapper.GetReleaseBranchNames(pkgConfig),
			}
		}
		cli_out.DumpPkgsTable(pkgInfos)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgListCommand)
}
