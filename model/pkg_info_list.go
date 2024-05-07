package model

import (
	"log"
)

// PkgInfoList - just a container for a list of PkgInfo with convenience search functionality
type PkgInfoList struct {
	PkgInfos []PkgInfo
}

// FindOneByNameOrFail - find a package by name
func (rcl *PkgInfoList) FindOneByNameOrFail(name string) *PkgInfo {
	for _, pkgInfo := range rcl.PkgInfos {
		if pkgInfo.PkgConfig.Name == name {
			return &pkgInfo
		}
	}
	log.Fatalln("Package not found: " + name)
	return nil
}
