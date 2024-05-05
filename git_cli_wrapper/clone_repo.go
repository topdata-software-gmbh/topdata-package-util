package git_cli_wrapper

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
)

func CloneRepo(repoConfig model.PkgConfig) {
	// Execute the git command to clone the repository
	folderName := repoConfig.GetLocalGitRepoDir()

	_ = execCommand("git", "clone", repoConfig.URL, folderName)
}

// RefreshRepo .. aka DownsyncRepo .... pulls all remote branches and checks them out locally
func RefreshRepo(repoConfig model.PkgConfig) {
	color.Blue(">>>> Refreshing repo: %s", repoConfig.Name)
	// check if repo exists
	if !repoConfig.IsLocalRepoExisting() {
		color.Yellow(">> Cloning repo: %s", repoConfig.Name)
		CloneRepo(repoConfig)
	}

	// fetch all remote branches
	remoteBranchNames := GetRemoteBranchNames(repoConfig)
	localBranchNames := GetLocalBranchNames(repoConfig)

	color.Cyan(">>>>>>>> Local branches: %v", localBranchNames)
	color.Cyan(">>>>>>>> Remote branches: %v", remoteBranchNames)

	// check out each remote branch locally if it doesn't already exist
	for _, branchName := range remoteBranchNames {
		if util.StringSliceContains(localBranchNames, branchName) {
			_ = execGitCommand(repoConfig, "checkout", branchName)
		} else {
			_ = execGitCommand(repoConfig, "checkout", "-b", branchName, "origin/"+branchName)
		}
	}
	// TODO: remove stale local branches
}
