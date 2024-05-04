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

func GetBranchNames(repoConf model.PkgConfig) []string {
	fmt.Println(">>>> GetBranchNames: " + repoConf.Name)
	// pkg for-each-ref --format='%(refname:short)' refs/heads/
	out, err := execGitCommand(repoConf, "for-each-ref", "--format", "%(refname:short)", "refs/heads/")
	if err != nil {
		log.Fatalln("Error getting branch names: " + err.Error())
	}
	log.Println("Branch names: " + out)

	branches := strings.Split(strings.TrimSpace(out), "\n")

	return branches
}

func GetReleaseBranchNames(repoConfig model.PkgConfig) []string {
	branchNames := GetBranchNames(repoConfig)
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
	out, _ := execGitCommand(repoConfig, "rev-parse", "HEAD")
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
	_, _ = execGitCommand(repoConfig, "checkout", branchName)
	_, _ = execGitCommand(repoConfig, "pull")
}

func GetReleaseBranches(repoConfig model.PkgConfig) []model.GitBranchInfo {
	releaseBranchNames := GetBranchNames(repoConfig)
	color.Yellow("Release branch names: %v\n", releaseBranchNames)
	releaseBranches := make([]model.GitBranchInfo, len(releaseBranchNames))
	for idx, branchName := range releaseBranchNames {
		fmt.Println("-----> " + branchName)
		releaseBranches[idx] = GetOneBranch(repoConfig, branchName)
	}
	// color.Blue("Release branches: %v\n", releaseBranches)
	return releaseBranches
}

func SwitchBranch(repoConfig model.PkgConfig, branchName string) error {
	_, err := execGitCommand(repoConfig, "checkout", branchName)
	if err != nil {
		log.Fatalln("Error switching branch: " + err.Error())
	}
	return nil
}
