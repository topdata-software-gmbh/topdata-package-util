package pkg_commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-util/factory"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"github.com/topdata-software-gmbh/topdata-package-util/printer"
)

// TODO.... for now it only shows table with single row .. fix that and show more details in a definition list like manner
var showGitBranchDetailsCommand = &cobra.Command{
	Use:   "show-git-branch-details [packageName] [branchName]",
	Short: "Shows details of single branch of a repository",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Details for repository: %s, branch %s ...\n", args[0], args[1])

		// ---- args
		pkgConfig := factory.NewPkgConfig(args[0])
		gitBranchInfo := factory.NewGitBranchInfo(pkgConfig, args[1])

		// ---- other info
		//pkgInfo := factory.NewPkgInfo(*pkgConfig)
		dict := map[string]string{
			"MachineName": pkgConfig.Name,
			"Git URL":     pkgConfig.URL,
		}
		printer.DumpDefinitionList(dict)

		// ---- table with branches
		fmt.Printf("Branch details: %v\n", gitBranchInfo)
		printer.DumpGitBranchInfoList(model.GitBranchInfoList{GitBranchInfos: []model.GitBranchInfo{gitBranchInfo}})

	},
}

func init() {
	pkgRootCommand.AddCommand(showGitBranchDetailsCommand)
}
