package git_service_v2

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"os/exec"
	"strings"
)

func FetchRepoBranches(repoURL string) ([]string, error) {
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
