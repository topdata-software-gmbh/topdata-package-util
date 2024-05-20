package model

import (
	"log"
	"sort"
)

// PkgInfoList - just a container for a list of PkgInfo with convenience search functionality
type PkgInfoList struct {
	PkgConfigList *PkgConfigList // the source of the list
	PkgInfos      []PkgInfo
	IsFiltered    bool
}

// FindOneByNameOrFail - find a package by name
func (list *PkgInfoList) FindOneByNameOrFail(name string) *PkgInfo {
	for _, pkgInfo := range list.PkgInfos {
		if pkgInfo.PkgConfig.Name == name {
			return &pkgInfo
		}
	}
	log.Fatalln("Package not found: " + name)
	return nil
}

func (list *PkgInfoList) FilterInShopware6Store() *PkgInfoList {
	filtered := PkgInfoList{
		PkgConfigList: list.PkgConfigList,
		IsFiltered:    true,
	}
	for _, pkgInfo := range list.PkgInfos {
		if pkgInfo.PkgConfig.InShopware6Store {
			filtered.PkgInfos = append(filtered.PkgInfos, pkgInfo)
		}
	}
	return &filtered
}

// Sort - sorts the PkgInfos slice based on the provided comparison function
func (list *PkgInfoList) Sort(compare func(a, b PkgInfo) bool) {
	sort.Slice(list.PkgInfos, func(i, j int) bool {
		return compare(list.PkgInfos[i], list.PkgInfos[j])
	})
}

// SortByName - sorts the PkgInfos slice by the Name field of PkgConfig
func (list *PkgInfoList) SortByName() {
	list.Sort(func(a, b PkgInfo) bool {
		return a.PkgConfig.Name < b.PkgConfig.Name
	})
}
