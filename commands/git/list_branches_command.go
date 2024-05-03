package git

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/cli_out"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
)

var listBranchesCommand = &cobra.Command{
	Use:   "list-branches [repositoryName]",
	Short: "Prints a table with all branches of a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s ...\n", args[0])

		repoConfig := model.GitRepoConfig{Name: args[0]}
		branchInfos := git_cli_wrapper.GetReleaseBranches(repoConfig)

		cli_out.DumpBranchesTable(branchInfos)
	},
}

func init() {
	gitRootCmd.AddCommand(listBranchesCommand)
}
