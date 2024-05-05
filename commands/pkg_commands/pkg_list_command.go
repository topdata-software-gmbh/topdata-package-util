package pkg_commands

import (
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
)

var pkgListCommand = &cobra.Command{
	Use:   "list",
	Short: "Prints a table with all packages",
	Run: func(cmd *cobra.Command, args []string) {
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)

		pkgInfos := make([]model.PkgInfo, len(pkgConfigList.PkgConfigs))
		for i, pkgConfig := range pkgConfigList.PkgConfigs {
			git_cli_wrapper.DownsyncRepo(pkgConfig)
			pkgInfos[i] = model.PkgInfo{
				Name:               pkgConfig.Name,
				URL:                pkgConfig.URL,
				BranchNames:        git_cli_wrapper.GetLocalBranchNames(pkgConfig),
				ReleaseBranchNames: git_cli_wrapper.GetReleaseBranchNames(pkgConfig),
			}
		}
		cli_out.DumpPkgsTable(pkgInfos)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgListCommand)
}
