package printer

import (
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
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
		"Commit Message",
	})

	for _, b := range gitBranchInfoList.GitBranchInfos {
		t.AppendRow([]interface{}{
			b.Name,
			b.PackageVersion,
			b.ShopwareVersionConstraint,
			b.CommitIdShort,
			b.CommitDate,
			b.CommitAuthor,
			b.CommitMessage,
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
		t.AppendHeader(table.Row{
			"In SW6 Store",
			"Package Name",
			"Release Branch Names",
			"Other Branch Names",
			// "URL"
		})
	} else {
		t.AppendHeader(table.Row{
			"Package Name",
			"Release Branch Names",
		})
	}
	for _, p := range pkgInfoList.PkgInfos {
		if displayMode == "full" {
			t.AppendRow([]interface{}{
				util.FormatBool(p.PkgConfig.InShopware6Store, "yes", ""),
				p.PkgConfig.Shopware6StoreTechnicalName,
				p.PkgConfig.Name,
				p.ReleaseBranchNames,
				p.OtherBranchNames,
				// p.URL
			})
		} else {
			t.AppendRow([]interface{}{
				p.PkgConfig.Name,
				p.ReleaseBranchNames,
			})
		}
	}

	t.Render()
}

func DumpDefinitionList(definitions map[string]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	//t.AppendHeader(table.Row{"Key", "Value"})

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
		"Commit Message",
	})

	for _, p := range ret {
		var branchName string
		var shopwareVersionConstraint string
		var commitId string
		var commitDate string
		var commitAuthor string
		var commitMessage string

		if p.Branch != nil {
			branchName = p.Branch.Name
			shopwareVersionConstraint = p.Branch.ShopwareVersionConstraint
			commitId = p.Branch.CommitIdShort
			commitDate = p.Branch.CommitDate
			commitAuthor = p.Branch.CommitAuthor
			commitMessage = p.Branch.CommitMessage
		} else {
			const PLACEHOLDER_NO_BRANCH_FOUND = "---"
			branchName = PLACEHOLDER_NO_BRANCH_FOUND
			shopwareVersionConstraint = PLACEHOLDER_NO_BRANCH_FOUND
			commitId = PLACEHOLDER_NO_BRANCH_FOUND
			commitDate = PLACEHOLDER_NO_BRANCH_FOUND
			commitAuthor = PLACEHOLDER_NO_BRANCH_FOUND
			commitMessage = PLACEHOLDER_NO_BRANCH_FOUND
		}

		t.AppendRow([]interface{}{
			p.Pkg.PkgConfig.Name,
			branchName,
			shopwareVersionConstraint,
			commitId,
			commitDate,
			commitAuthor,
			commitMessage,
		})
	}

	t.Render()
}

func DumpBranchesDiffTable(branchNames []string, branchesDiff []map[string]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	// ---- header
	headerRow := table.Row{
		"commitId",
		"date",
		//		"author",
		"message",
	}
	// append the branchNames to the header row
	for _, branchName := range branchNames {
		headerRow = append(headerRow, branchName)
	}
	t.AppendHeader(headerRow)

	// ---- rows
	for _, commit := range branchesDiff {
		row := table.Row{
			commit["commitId"],
			commit["date"],
			//			commit["author"],
			commit["message"],
		}
		// append the branchNames to the row
		for _, branchName := range branchNames {
			row = append(row, commit[branchName])
		}
		t.AppendRow(row)
	}

	t.Render()
}
