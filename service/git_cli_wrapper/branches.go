package git_cli_wrapper

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"regexp"
	"strings"
)

func GetBranchNames(repoConf model.GitRepoConfig) []string {
	fmt.Println(">>>> GetBranchNames: " + repoConf.Name)
	// git for-each-ref --format='%(refname:short)' refs/heads/
	out, err := execGitCommand(repoConf, "for-each-ref", "--format", "%(refname:short)", "refs/heads/")
	if err != nil {
		log.Fatalln("Error getting branch names: " + err.Error())
	}
	log.Println("Branch names: " + out)

	branches := strings.Split(strings.TrimSpace(out), "\n")

	return branches
}

func GetReleaseBranchNames(repoConfig model.GitRepoConfig) []string {
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

func getBranchCommitId(repoConfig model.GitRepoConfig, name2 string) string {
	// git rev-parse refs/heads/branchName

	out, err := execGitCommand(repoConfig, "rev-parse", "refs/heads/"+name2)
	if err != nil {
		log.Fatalln("Error getting branch commit id: " + err.Error())
	}

	return strings.TrimSpace(out)
}

func GetOneBranch(repoName string, branchName string) model.GitBranchInfo {
	branchInfo := model.GitBranchInfo{
		Name:     branchName,
		CommitId: getBranchCommitId(model.GitRepoConfig{Name: repoName}, branchName),
	}
	fmt.Println("Branch details for repository: " + repoName + ", branch: " + branchName)

	return branchInfo
}

func GetReleaseBranches(repositoryName string) []model.GitBranchInfo {
	releaseBranchNames := GetBranchNames(model.GitRepoConfig{Name: repositoryName})
	color.Yellow("Release branch names: %v\n", releaseBranchNames)
	releaseBranches := make([]model.GitBranchInfo, len(releaseBranchNames))
	for idx, branchName := range releaseBranchNames {
		fmt.Println("-----> " + branchName)
		releaseBranches[idx] = GetOneBranch(repositoryName, branchName)
	}
	color.Blue("Release branches: %v\n", releaseBranches)
	return releaseBranches
}
