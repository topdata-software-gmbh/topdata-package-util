package model

import (
	"github.com/Masterminds/semver/v3"
	"log"
	"sort"
)

type GitBranchInfoList struct {
	GitBranchInfos []GitBranchInfo
}

func (g GitBranchInfoList) Len() int {
	return len(g.GitBranchInfos)
}

func (g GitBranchInfoList) SortByPackageVersionAsc() {
	sort.Slice(g.GitBranchInfos, func(i, j int) bool {
		vi, err := semver.NewVersion(g.GitBranchInfos[i].PackageVersion)
		if err != nil {
			log.Fatalln("Error parsing version: %s", err)
		}

		vj, err := semver.NewVersion(g.GitBranchInfos[j].PackageVersion)
		if err != nil {
			log.Fatalln("Error parsing version: %s", err)
		}

		return vi.LessThan(vj)
	})
}
