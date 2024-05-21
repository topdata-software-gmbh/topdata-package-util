package localgit_commands

import (
	"github.com/spf13/cobra"
)

// TODO: finish this command

var compareBranchesCommand = &cobra.Command{
	Use:   "compare-branches [branch1] [branch2]",
	Short: "Compares two branches and show the differences in a table",
	Args:  cobra.ExactArgs(2), // expects exactly 2 arguments
	RunE: func(cmd *cobra.Command, args []string) error {
		// args[0] will be the first argument, and args[1] the second
		branch1 := args[0]
		branch2 := args[1]

		// print the branch names
		cmd.Printf("TODO: Comparing branches %s and %s in CWD repo\n", branch1, branch2)

		return nil
	},
}

func init() {
	localgitRootCommand.AddCommand(compareBranchesCommand)
}
