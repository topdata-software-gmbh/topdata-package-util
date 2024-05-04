package cli_out

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
)

func DumpBranchesTable(branches []model.GitBranchInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Git Branch", "Commit Id", "Package Version", "Shopware Version"})

	for _, b := range branches {
		t.AppendRow([]interface{}{b.Name, b.CommitId, b.PackageVersion, b.ShopwareVersion})
	}

	t.Render()
}

func DumpPkgsTable(pkgInfos []model.PkgInfo) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Package Name", "URL", "Branch Names", "Release Branch Names"})

	for _, p := range pkgInfos {
		t.AppendRow([]interface{}{p.Name, p.URL, p.BranchNames, p.ReleaseBranchNames})
	}

	t.Render()
}
