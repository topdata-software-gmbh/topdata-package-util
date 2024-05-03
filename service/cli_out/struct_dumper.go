package cli_out

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
)

func DumpBranchesTable(branches []model.GitBranchInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Commit Id", "Version"})

	for _, b := range branches {
		t.AppendRow([]interface{}{b.Name, b.CommitId, b.Version})
	}

	t.Render()
}
