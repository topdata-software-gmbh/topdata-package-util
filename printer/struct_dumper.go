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
	t.AppendHeader(table.Row{
		"Git Branch",
		"Package Version",
		"Shopware Version",
		"Commit Id",
		"Commit Date",
		"Commit Author",
	})

	for _, b := range gitBranchInfoList.GitBranchInfos {
		t.AppendRow([]interface{}{
			b.Name,
			b.PackageVersion,
			b.ShopwareVersionConstraint,
			b.CommitId,
			b.CommitDate,
			b.CommitAuthor,
		})
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
			t.AppendRow([]interface{}{p.PkgConfig.Name, p.ReleaseBranchNames, p.OtherBranchNames /*, p.URL*/})
		} else {
			t.AppendRow([]interface{}{p.PkgConfig.Name, p.ReleaseBranchNames})
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

func DumpPkgAndBranchTable(ret []model.PkgAndBranch) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Package",
		"Branch",
		"Shopware Version",
		"Commit Id",
		"Commit Date",
		"Commit Author",
	})

	for _, p := range ret {
		var branchName string
		var shopwareVersionConstraint string
		var commitId string
		var commitDate string
		var commitAuthor string

		if p.Branch != nil {
			branchName = p.Branch.Name
			shopwareVersionConstraint = p.Branch.ShopwareVersionConstraint
			commitId = p.Branch.CommitId
			commitDate = p.Branch.CommitDate
			commitAuthor = p.Branch.CommitAuthor
		} else {
			const PLACEHOLDER_NO_BRANCH_FOUND = "---"
			branchName = PLACEHOLDER_NO_BRANCH_FOUND
			shopwareVersionConstraint = PLACEHOLDER_NO_BRANCH_FOUND
			commitId = PLACEHOLDER_NO_BRANCH_FOUND
			commitDate = PLACEHOLDER_NO_BRANCH_FOUND
			commitAuthor = PLACEHOLDER_NO_BRANCH_FOUND
		}

		t.AppendRow([]interface{}{
			p.Pkg.PkgConfig.Name,
			branchName,
			shopwareVersionConstraint,
			commitId,
			commitDate,
			commitAuthor,
		})
	}

	t.Render()
}
