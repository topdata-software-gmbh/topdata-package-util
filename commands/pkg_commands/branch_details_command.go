package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
)

// TODO.... for now it only shows table with single row .. fix that and show more details in a definition list like manner
var showGitBranchDetailsCommand = &cobra.Command{
	Use:   "show-git-branch-details [packageName] [branchName]",
	Short: "Shows details of single branch of a repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s, branch %s ...\n", args[0], args[1])

		repoConfig := model.PkgConfig{Name: args[0]}
		gitBranchInfo := git_cli_wrapper.GetOneBranch(repoConfig, args[1])
		fmt.Printf("Branch details: %v\n", gitBranchInfo)
		cli_out.DumpGitBranchInfoList(model.GitBranchInfoList{GitBranchInfos: []model.GitBranchInfo{gitBranchInfo}})
	},
}

func init() {
	pkgRootCommand.AddCommand(showGitBranchDetailsCommand)
}
