package git_service_v2

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	_ "github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
)

func GetBranchDetails(repoName string, branchName string) model.GitBranchInfo {
	branchInfo := model.GitBranchInfo{
		Name: branchName,
	}
	fmt.Println("Branch details for repository: " + repoName + ", branch: " + branchName)

	return branchInfo
}
