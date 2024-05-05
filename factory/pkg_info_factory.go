package factory

import (
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
)

// NewPkgInfo creates a new PkgInfo object (aka constructor)
func NewPkgInfo(pkgConfig model.PkgConfig) model.PkgInfo {
	git_cli_wrapper.DownsyncRepo(pkgConfig)
	branchNames := git_cli_wrapper.GetRemoteBranchNames(pkgConfig)
	regex := `^(main|main-.*|release-.*)$` // TODO: the regex should be part of the service config or even pkgConfig

	return model.PkgInfo{
		Name:               pkgConfig.Name,
		URL:                pkgConfig.URL,
		ReleaseBranchNames: util.FilterStringArrayPositive(branchNames, regex),
		OtherBranchNames:   util.FilterStringArrayNegative(branchNames, regex),
	}

}
