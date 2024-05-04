package git_cli_wrapper

// This service uses the pkg CLI to interact with pkg repositories.
// It is a replacement for the go-pkg library.

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"regexp"
	"strings"
)

func GetLocalBranchNames(repoConf model.PkgConfig) []string {
	fmt.Println(">>>> GetLocalBranchNames: " + repoConf.Name)
	// git for-each-ref --format='%(refname:short)' refs/heads/
	out := execGitCommand(repoConf, "for-each-ref", "--format", "%(refname:short)", "refs/heads/")
	branches := strings.Split(strings.TrimSpace(out), "\n")

	return branches
}

func GetRemoteBranchNames(pkgConfig model.PkgConfig) []string {
	// git ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'
	//out := execCommand("sh", "-c", "git ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'")

	shellCommand := fmt.Sprintf("GIT_SSH_COMMAND='ssh -i %s' git -C %s ls-remote --heads origin | awk '{print $2}' | sed 's#refs/heads/##' | sed '/^$/d'", *pkgConfig.PathSshKey, pkgConfig.GetLocalGitRepoDir())
	fmt.Println("================= cmd: ", shellCommand)
	out := execShellCommand(shellCommand)

	return strings.Split(strings.TrimSpace(out), "\n")
}

func GetReleaseBranchNames(repoConfig model.PkgConfig) []string {
	branchNames := GetLocalBranchNames(repoConfig)
	return filterBranchNames(branchNames, `^(main|main-.*|release-.*)$`)
}

// filterBranchNames filters the given branches and returns only those that are either "server" or start with "release-".
// It takes a slice of strings representing the branch names as input and returns a slice of strings containing the filtered branch names.
func filterBranchNames(branches []string, regexPattern string) []string {
	releaseBranches := make([]string, 0)
	for _, branch := range branches {
		// TODO: the regex should be part of the service config
		matched, _ := regexp.MatchString(regexPattern, branch)
		if matched {
			releaseBranches = append(releaseBranches, branch)
		}
	}
	return releaseBranches
}

// returns the commit id of the current branch
func getCommitId(repoConfig model.PkgConfig) string {
	out := execGitCommand(repoConfig, "rev-parse", "HEAD")
	return strings.TrimSpace(out)
}

//func getCommitId(repoConfig model.PkgConfig, name2 string) string {
//	// pkg rev-parse refs/heads/branchName
//
//	out, err := execGitCommand(repoConfig, "rev-parse", "refs/heads/"+name2)
//	if err != nil {
//		log.Fatalln("Error getting branch commit id: " + err.Error())
//	}
//
//	return strings.TrimSpace(out)
//}

func GetOneBranch(repoConfig model.PkgConfig, branchName string) model.GitBranchInfo {

	checkoutBranch(repoConfig, branchName)

	composerJson := getComposerJson(repoConfig)
	branchInfo := model.GitBranchInfo{
		Name:            branchName,
		CommitId:        getCommitId(repoConfig),
		PackageVersion:  composerJson.Version,
		ShopwareVersion: composerJson.Require["shopware/core"],
	}
	fmt.Println("Branch details for repository: " + repoConfig.Name + ", branch: " + branchName)

	return branchInfo
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

func checkoutBranch(repoConfig model.PkgConfig, branchName string) {
	_ = execGitCommand(repoConfig, "checkout", branchName)
	_ = execGitCommand(repoConfig, "pull")
}

func GetReleaseBranches(repoConfig model.PkgConfig) []model.GitBranchInfo {
	releaseBranchNames := GetLocalBranchNames(repoConfig)
	color.Yellow("Release branch names: %v\n", releaseBranchNames)
	releaseBranches := make([]model.GitBranchInfo, len(releaseBranchNames))
	for idx, branchName := range releaseBranchNames {
		fmt.Println("-----> " + branchName)
		releaseBranches[idx] = GetOneBranch(repoConfig, branchName)
	}
	// color.Blue("Release branches: %v\n", releaseBranches)
	return releaseBranches
}

func SwitchBranch(pkgConfig model.PkgConfig, branchName string) {
	log.Println("Switching to branch: " + branchName)
	_ = execGitCommand(pkgConfig, "checkout", branchName)
}
