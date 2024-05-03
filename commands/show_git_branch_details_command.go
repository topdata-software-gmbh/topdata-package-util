package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_service_v2"
)

var showGitBranchDetailsCommand = &cobra.Command{
	Use:   "show-git-branch-details [repositoryName] [branchName]",
	Short: "Testing git cli wrapper",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s, branch %s ...\n", args[0], args[1])
		gitBranchInfo := git_service_v2.GetBranchDetails(args[0], args[1])
		fmt.Printf("Branch details: %v\n", gitBranchInfo)
		git_service_v2.PrintBranchesTable([]model.GitBranchInfo{gitBranchInfo})
	},
}

func init() {
	rootCmd.AddCommand(showGitBranchDetailsCommand)
}
