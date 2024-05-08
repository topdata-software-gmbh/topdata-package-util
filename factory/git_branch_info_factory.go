package factory

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
	"log"
)

// NewGitBranchInfo creates a new GitBranchInfo object (aka constructor)
func NewGitBranchInfo(pkgConfig model.PkgConfig, branchName string) model.GitBranchInfo {
	git_cli_wrapper.CheckoutBranch(pkgConfig, branchName)

	composerJson := getComposerJson(pkgConfig)
	branchInfo := model.GitBranchInfo{
		Name:                      branchName,
		CommitId:                  git_cli_wrapper.GetCommitId(pkgConfig),
		CommitDate:                git_cli_wrapper.GetCommitDate(pkgConfig),
		CommitAuthor:              git_cli_wrapper.GetCommitAuthor(pkgConfig),
		PackageVersion:            composerJson.Version,
		ShopwareVersionConstraint: composerJson.Require["shopware/core"],
	}
	return branchInfo
}

// NewBranchInfos creates multiple GitBranchInfo objects, one for each release branch
func NewBranchInfos(repoConfig model.PkgConfig, onlyReleaseBranches bool) []model.GitBranchInfo {
	releaseBranchNames := git_cli_wrapper.GetLocalBranchNames(repoConfig)
	if onlyReleaseBranches {
		releaseBranchNames = util.FilterStringSlicePositive(releaseBranchNames, `^(main|main-.*|release-.*)$`)
	}

	color.Yellow("Release branch names: %v\n", releaseBranchNames)
	releaseBranches := make([]model.GitBranchInfo, len(releaseBranchNames))
	for idx, branchName := range releaseBranchNames {
		// fmt.Println("-----> " + branchName)
		releaseBranches[idx] = NewGitBranchInfo(repoConfig, branchName)
	}
	// color.Blue("Release branches: %v\n", releaseBranches)
	return releaseBranches
}

func getComposerJson(repoConfig model.PkgConfig) model.ComposerJSON {
	var composerJson model.ComposerJSON

	// Load data from composer.json file
	err := composerJson.LoadFromFile(repoConfig.GetAbsolutePath("composer.json"))
	if err != nil {
		log.Fatalln("Error loading composer.json: " + err.Error())
	}

	return composerJson
}
