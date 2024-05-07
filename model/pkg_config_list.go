package model

import "log"

// PkgConfigList - just a container for a list of PkgConfig with convenience search functionality
type PkgConfigList struct {
	MachineName string // used for cache file name
	PkgConfigs  []PkgConfig
}

func (rcl *PkgConfigList) FindOneByNameOrFail(name string) *PkgConfig {
	for _, pkgConfig := range rcl.PkgConfigs {
		if pkgConfig.Name == name {
			return &pkgConfig
		}
	}
	log.Fatalln("Package not found: " + name)
	return nil
}
