package git_cli_wrapper

// This service uses the pkg CLI to interact with pkg repositories.
// It is a replacement for the go-pkg library.

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
	"strings"
)

// GetLocalBranchNames - git for-each-ref --format='%(refname:short)' refs/heads/
func GetLocalBranchNames(pkgConfig *model.PkgConfig) []string {
	color.Blue(">>>> GetLocalBranchNames: " + pkgConfig.Name)
	out := runGitCommandInClonedRepo(pkgConfig, "for-each-ref", "--format", "%(refname:short)", "refs/heads/")
	branches := strings.Split(strings.TrimSpace(out), "\n")

	return branches
}

// GetRemoteBranchNames - git ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'
func GetRemoteBranchNames(pkgConfig *model.PkgConfig) []string {
	color.Blue(">>>> GetRemoteBranchNames: " + pkgConfig.Name)
	// extra env to set ssh key
	var extraEnv []string
	if pkgConfig.PathSshKey != nil {
		extraEnv = append(extraEnv, fmt.Sprintf("GIT_SSH_COMMAND=/usr/bin/ssh -i %s", *pkgConfig.PathSshKey))
	}

	// run shell command
	shellCommand := fmt.Sprintf("git -C %s ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'", pkgConfig.GetLocalGitRepoDir())
	out := util.RunShellCommand(shellCommand, &extraEnv)

	return strings.Split(strings.TrimSpace(out), "\n")
}

// returns the commit id of the current branch
func GetCommitId(pkgConfig *model.PkgConfig) string {
	out := runGitCommandInClonedRepo(pkgConfig, "rev-parse", "HEAD")
	return strings.TrimSpace(out)
}

func GetCommitIdShort(pkgConfig *model.PkgConfig) string {
	out := runGitCommandInClonedRepo(pkgConfig, "rev-parse", "--short", "HEAD")
	return strings.TrimSpace(out)
}

func GetCommitMessage(pkgConfig *model.PkgConfig) string {
	out := runGitCommandInClonedRepo(pkgConfig, "show", "-s", "--format=%s", "HEAD")
	return strings.TrimSpace(out)
}

func GetCommitDate(pkgConfig *model.PkgConfig) string {
	out := runGitCommandInClonedRepo(pkgConfig, "show", "-s", "--format=%ci", "HEAD")
	return strings.TrimSpace(out)
}

func GetCommitAuthor(pkgConfig *model.PkgConfig) string {
	out := runGitCommandInClonedRepo(pkgConfig, "show", "-s", "--format=%an", "HEAD")
	return strings.TrimSpace(out)
}

//func GetCommitId(pkgConfig *model.PkgConfig, name2 string) string {
//	// pkg rev-parse refs/heads/branchName
//
//	out, err := runGitCommandInClonedRepo(pkgConfig, "rev-parse", "refs/heads/"+name2)
//	if err != nil {
//		log.Fatalln("Error getting branch commit id: " + err.Error())
//	}
//
//	return strings.TrimSpace(out)
//}

func CheckoutBranch(pkgConfig *model.PkgConfig, branchName string) {
	_ = runGitCommandInClonedRepo(pkgConfig, "checkout", "-f", branchName)
	_ = runGitCommandInClonedRepo(pkgConfig, "pull")
}

func SwitchBranch(pkgConfig *model.PkgConfig, branchName string) {
	_ = runGitCommandInClonedRepo(pkgConfig, "checkout", "-f", branchName)
}

func CompareBranches(pkgConfig *model.PkgConfig) {
	// Get all unique commits from both branches
	gitCmd := "git -C " + pkgConfig.GetLocalGitRepoDir() + " log --date=format:'%Y-%m-%d %H:%M:%S' --pretty=format:'%h,%ad,%an,%s' --all | sort | uniq | sort -t',' -k2,2r"
	println(gitCmd)
	output := util.RunShellCommand(gitCmd, nil)
	fmt.Println(output)
}
