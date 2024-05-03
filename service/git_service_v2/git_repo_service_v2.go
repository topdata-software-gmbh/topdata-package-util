package git_service_v2

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/file_path_service"
	"os/exec"
	"strings"
)

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

func FetchRepositoryBranches(repoURL string) ([]string, error) {
	// Execute the git command to fetch all branches
	cmd := exec.Command("git", "ls-remote", "--heads", repoURL)
	output, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	// Parse the output to get the branch names
	branches := strings.Split(string(output), "\n")
	for i, branch := range branches {
		branches[i] = strings.TrimPrefix(branch, "refs/heads/")
	}

	return branches, nil
}

func CloneRepository(repoConfig model.GitRepositoryConfig) error {
	// Execute the git command to clone the repository
	folderName := file_path_service.GetLocalGitRepoDir(repoConfig)

	err2 := execCommand("git", "clone", repoConfig.URL, folderName)
	if err2 != nil {
		return err2
	}

	return nil
}

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))

	if err != nil {
		color.Yellow("!!!!! cmd: " + cmd.String())
		color.Red("!!!!! out: " + strings.TrimSpace(string(output)))
		color.Yellow("!!!!! err: " + err.Error())

		return err
	}
	return nil
}
