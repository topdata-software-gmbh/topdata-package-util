package pkg_commands

import (
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-service/globals"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/printer"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
)

var bShowAllBranches bool

var pkgDetailsCommand = &cobra.Command{
	Use:   "details [packageName]",
	Short: "Prints a table with all branches of a repository and some other info",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Details for repository: %s ...\n", args[0])

		// ---- load the package portfolio file
		// pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		// pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)
		pkgConfig := globals.PkgConfigList.FindOneByNameOrFail(args[0])
		git_cli_wrapper.RefreshRepo(*pkgConfig)

		// ---- other info
		//pkgInfo := factory.NewPkgInfo(*pkgConfig)
		dict := map[string]string{
			"MachineName":        pkgConfig.Name,
			"Local Repo":         pkgConfig.GetLocalGitRepoDir(),
			"Git URL":            pkgConfig.URL,
			"In Shopware6 Store": util.FormatBool(pkgConfig.InShopware6Store, "yes", ""),
		}
		printer.DumpDefinitionList(dict)

		// ---- branches in a table
		branchInfoList := factory.NewBranchInfos(*pkgConfig, !bShowAllBranches)
		printer.DumpGitBranchInfoList(model.GitBranchInfoList{GitBranchInfos: branchInfoList})
	},
}

func init() {
	pkgDetailsCommand.Flags().BoolVarP(&bShowAllBranches, "all", "a", false, "Show all branches (not only release branches)")
	pkgRootCommand.AddCommand(pkgDetailsCommand)
}
