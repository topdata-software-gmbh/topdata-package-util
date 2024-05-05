package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/config"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
)

var listBranchesCommand = &cobra.Command{
	Use:   "details [packageName]",
	Short: "Prints a table with all branches of a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s ...\n", args[0])

		// ---- load the package portfolio file
		pathPackagesPortfolioFile, _ := cmd.Flags().GetString("packages-portfolio-file")
		pkgConfigList := config.LoadPackagePortfolioFile(pathPackagesPortfolioFile)
		pkgConfig := pkgConfigList.FindOneByNameOrFail(args[0])

		//git_cli_wrapper.CloneRepo(*pkgConfig) // TODO: CloneIfNotExists

		branchNames := git_cli_wrapper.GetLocalBranchNames(*pkgConfig)
		fmt.Printf("+++++++++++++++ Branches: %v\n", branchNames)

		remoteBranchNames := git_cli_wrapper.GetRemoteBranchNames(*pkgConfig)
		fmt.Printf("+++++++++++++++ Remote Branches: %v\n", remoteBranchNames)

		// ----
		branchInfoList := git_cli_wrapper.GetReleaseBranches(*pkgConfig)

		cli_out.DumpGitBranchInfoList(model.GitBranchInfoList{GitBranchInfos: branchInfoList})
	},
}

func init() {
	pkgRootCommand.AddCommand(listBranchesCommand)
}
