package cli_out

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
)

func DumpBranchesTable(branches []model.GitBranchInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Git Branch", "Commit Id", "PackageVersion", "Shopware PackageVersion"})

	for _, b := range branches {
		t.AppendRow([]interface{}{b.Name, b.CommitId, b.PackageVersion, b.ShopwareVersion})
	}

	t.Render()
}
