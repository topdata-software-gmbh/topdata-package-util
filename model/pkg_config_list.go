package model

import "log"

// aka portfolio
// PkgConfigList - just a container for a list of PkgConfig with convenience search functionality
type PkgConfigList struct {
	MachineName string      `yaml:"machineName"` // used for cache file name
	Items       []PkgConfig `yaml:"items"`
}

func (rcl *PkgConfigList) FindOneByNameOrFail(name string) *PkgConfig {
	for _, pkgConfig := range rcl.Items {
		if pkgConfig.Name == name {
			return &pkgConfig
		}
	}
	log.Fatalln("Package not found: " + name)
	return nil
}
