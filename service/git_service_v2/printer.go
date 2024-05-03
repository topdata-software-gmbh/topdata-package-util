package git_service_v2

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
)

func PrintBranchesTable(branches []model.GitBranchInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "CommitId"})

	for _, b := range branches {
		t.AppendRow([]interface{}{b.Name, b.CommitId})
	}

	t.Render()
}
