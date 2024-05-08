package model

import (
	"log"
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
