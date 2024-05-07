package factory

import "github.com/topdata-software-gmbh/topdata-package-service/model"

func NewPkgConfig(Name string) model.PkgConfig {
	// FIXME: it should search the item in the loaded portfolio
	return model.PkgConfig{Name: Name}
}
