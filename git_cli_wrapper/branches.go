package git_cli_wrapper

// This service uses the pkg CLI to interact with pkg repositories.
// It is a replacement for the go-pkg library.

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
	"strings"
)

// GetLocalBranchNames - git for-each-ref --format='%(refname:short)' refs/heads/
func GetLocalBranchNames(repoConf model.PkgConfig) []string {
	color.Blue(">>>> GetLocalBranchNames: " + repoConf.Name)
	out := execGitCommand(repoConf, "for-each-ref", "--format", "%(refname:short)", "refs/heads/")
	branches := strings.Split(strings.TrimSpace(out), "\n")

	return branches
}

// GetRemoteBranchNames - git ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'
func GetRemoteBranchNames(pkgConfig model.PkgConfig) []string {
	color.Blue(">>>> GetRemoteBranchNames: " + pkgConfig.Name)
	var extraEnv []string
	if pkgConfig.PathSshKey != nil {
		extraEnv = append(extraEnv, fmt.Sprintf("GIT_SSH_COMMAND=/usr/bin/ssh -i %s", *pkgConfig.PathSshKey))
	}

	shellCommand := fmt.Sprintf("git -C %s ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'", pkgConfig.GetLocalGitRepoDir())
	out := util.ExecShellCommand(shellCommand, extraEnv)

	return strings.Split(strings.TrimSpace(out), "\n")
}

// returns the commit id of the current branch
func GetCommitId(repoConfig model.PkgConfig) string {
	out := execGitCommand(repoConfig, "rev-parse", "HEAD")
	return strings.TrimSpace(out)
}

//func GetCommitId(repoConfig model.PkgConfig, name2 string) string {
//	// pkg rev-parse refs/heads/branchName
//
//	out, err := execGitCommand(repoConfig, "rev-parse", "refs/heads/"+name2)
//	if err != nil {
//		log.Fatalln("Error getting branch commit id: " + err.Error())
//	}
//
//	return strings.TrimSpace(out)
//}

func CheckoutBranch(repoConfig model.PkgConfig, branchName string) {
	_ = execGitCommand(repoConfig, "checkout", branchName)
	_ = execGitCommand(repoConfig, "pull")
}

func SwitchBranch(pkgConfig model.PkgConfig, branchName string) {
	_ = execGitCommand(pkgConfig, "checkout", branchName)
}
