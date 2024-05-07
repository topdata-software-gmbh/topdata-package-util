package printer

import (
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
)

func DumpGitBranchInfoList(gitBranchInfoList model.GitBranchInfoList) {

	gitBranchInfoList.SortByPackageVersionAsc()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Git Branch", "Package Version", "Shopware Version", "Commit Id"})

	for _, b := range gitBranchInfoList.GitBranchInfos {
		t.AppendRow([]interface{}{b.Name, b.PackageVersion, b.ShopwareVersion, b.CommitId})
	}

	t.Render()
}

// DumpPkgInfoListTable dumps the PkgInfoList as a table to the console
func DumpPkgInfoListTable(pkgInfoList *model.PkgInfoList, displayMode string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	color.Blue("DumpPkgInfoListTable.displayMode=%s", displayMode)

	if displayMode == "full" {
		t.AppendHeader(table.Row{"Package Name", "Release Branch Names", "Other Branch Names" /*, "URL"*/})
	} else {
		t.AppendHeader(table.Row{"Package Name", "Release Branch Names"})
	}
	for _, p := range pkgInfoList.PkgInfos {
		if displayMode == "full" {
			t.AppendRow([]interface{}{p.Name, p.ReleaseBranchNames, p.OtherBranchNames /*, p.URL*/})
		} else {
			t.AppendRow([]interface{}{p.Name, p.ReleaseBranchNames})
		}
	}

	t.Render()
}

func DumpDefinitionList(definitions map[string]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Key", "Value"})

	for key, value := range definitions {
		t.AppendRow(table.Row{key, value})
	}

	t.Render()
}
