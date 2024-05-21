package git_cli_wrapper

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
)

func CloneRepo(pkgConfig *model.PkgConfig) {
	// Execute the git command to clone the repository
	folderName := pkgConfig.GetLocalGitRepoDir()

	_ = util.RunCommand("git", "clone", pkgConfig.URL, folderName)
}

// RefreshRepo .. aka DownsyncRepo .... pulls all remote branches and checks them out locally
func RefreshRepo(pkgConfig *model.PkgConfig) {
	color.Blue(">>>> Refreshing repo: %s", pkgConfig.Name)
	// check if repo exists
	if !pkgConfig.IsLocalRepoExisting() {
		color.Yellow(">> Cloning repo: %s", pkgConfig.Name)
		CloneRepo(pkgConfig)
	}

	// fetch all remote branches
	remoteBranchNames := GetRemoteBranchNames(pkgConfig)
	localBranchNames := GetLocalBranchNames(pkgConfig)

	color.Cyan(">>>>>>>> Local branches: %v", localBranchNames)
	color.Cyan(">>>>>>>> Remote branches: %v", remoteBranchNames)

	// check out each remote branch locally if it doesn't already exist
	for _, branchName := range remoteBranchNames {
		if util.StringSliceContains(localBranchNames, branchName) {
			_ = runGitCommandInClonedRepo(pkgConfig, "checkout", "-f", branchName)
		} else {
			_ = runGitCommandInClonedRepo(pkgConfig, "checkout", "-b", branchName, "origin/"+branchName)
		}
	}
	// TODO: remove stale local branches
}
