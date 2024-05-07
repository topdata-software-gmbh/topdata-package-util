package pkg_commands

import (
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/printer"
)

var pkgDetailsCommand = &cobra.Command{
	Use:   "details [packageName]",
	Short: "Prints a table with all branches of a repository and some other info",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Details for repository: %s ...\n", args[0])

		// ---- load the package portfolio file
		pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)
		pkgConfig := pkgConfigList.FindOneByNameOrFail(args[0])
		git_cli_wrapper.RefreshRepo(*pkgConfig)

		// ---- branches in a table
		branchInfoList := factory.NewReleaseBranchInfos(*pkgConfig)
		printer.DumpGitBranchInfoList(model.GitBranchInfoList{GitBranchInfos: branchInfoList})

		// ---- other info
		//pkgInfo := factory.NewPkgInfo(*pkgConfig)
		dict := map[string]string{
			"MachineName": pkgConfig.Name,
			"Local Dir":   pkgConfig.GetLocalGitRepoDir(),
			"Git URL":     pkgConfig.URL,
		}
		printer.DumpDefinitionList(dict)
	},
}

func init() {
	pkgRootCommand.AddCommand(pkgDetailsCommand)
}
