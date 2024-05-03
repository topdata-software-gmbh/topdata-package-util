package git_service_v2

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	_ "github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
)

func GetBranchDetails(repoName string, branchName string) model.GitBranchInfo {
	branchInfo := model.GitBranchInfo{
		Name:     branchName,
		CommitId: getBranchCommitId(model.GitRepoConfig{Name: repoName}, branchName),
	}
	fmt.Println("Branch details for repository: " + repoName + ", branch: " + branchName)

	return branchInfo
}

func getBranchCommitId(repoConfig model.GitRepoConfig, name2 string) string {
	// git rev-parse refs/heads/branchName

	out, err := execGitCommand(repoConfig, "rev-parse", "refs/heads/"+name2)
	if err != nil {
		log.Fatalln("Error getting branch commit id: " + err.Error())
	}

	return out
}
