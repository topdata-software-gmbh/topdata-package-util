package factory

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/app_constants"
	git_cli_wrapper2 "github.com/topdata-software-gmbh/topdata-package-util/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"github.com/topdata-software-gmbh/topdata-package-util/serializers"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
)

// NewPkgInfo creates a new PkgInfo object (aka constructor)
func NewPkgInfo(pkgConfig *model.PkgConfig) model.PkgInfo {

	color.Blue("//////////////// NewPkgInfo: %s", pkgConfig.Name)

	git_cli_wrapper2.RefreshRepo(pkgConfig)
	branchNames := git_cli_wrapper2.GetRemoteBranchNames(pkgConfig)
	regex := `^(main|main-.*|release-.*)$` // TODO: the regex should be part of the service config or even pkgConfig

	return model.PkgInfo{
		PkgConfig: pkgConfig,
		// Name:               pkgConfig.Name,
		// URL:                pkgConfig.URL,
		ReleaseBranchNames: util.FilterStringSlicePositive(branchNames, regex),
		OtherBranchNames:   util.FilterStringSliceNegative(branchNames, regex),
	}

}

func NewPkgInfoList(PkgConfigList *model.PkgConfigList) *model.PkgInfoList {
	pkgInfos := make([]model.PkgInfo, len(PkgConfigList.PkgConfigs))

	for i, pkgConfig := range PkgConfigList.PkgConfigs {
		pkgInfos[i] = NewPkgInfo(&pkgConfig)
	}

	return &model.PkgInfoList{
		PkgConfigList: PkgConfigList,
		IsFiltered:    false,
		PkgInfos:      pkgInfos,
	}
}

func NewPkgInfoListCached(PkgConfigList *model.PkgConfigList, bForceRefresh bool) *model.PkgInfoList {
	if util.FileExists(app_constants.PathCacheFile) && !bForceRefresh {
		color.Yellow(">>>> Loading from cache_commands file %s", app_constants.PathCacheFile)
		return serializers.LoadPkgInfoList(app_constants.PathCacheFile)
	} else {
		pkgInfoList := NewPkgInfoList(PkgConfigList)
		serializers.SavePkgInfoList(pkgInfoList, app_constants.PathCacheFile)
		return pkgInfoList
	}

}
