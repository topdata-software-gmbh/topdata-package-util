package git_repository_service

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"path/filepath"
	"regexp"
	"sort"
)

func GetRepositoryBranches(repoConf model.GitRepositoryConfig) ([]string, error) {
	fmt.Println(">>>> GetRepositoryBranches: " + repoConf.Name)
	// ---- fetch branches from the repoConf
	gitDir := filepath.Join("/tmp/git-repos", repoConf.Name)

	// ---- git clone / pull
	repo, err := refreshRepo(repoConf, gitDir)
	if err != nil {
		fmt.Println("Error fetching repo to " + gitDir + ": " + err.Error())
		return nil, err
	}

	remoteBranches, err := getRemoteBranches(repoConf, repo)
	if err != nil {
		fmt.Println("Error fetching remote branches to " + gitDir + ": " + err.Error())
		return nil, err
	}

	return remoteBranches, nil
}

func getBranches(repo *git.Repository) ([]string, error) {
	// ---- branches from the local repository
	branches := make([]string, 0)
	refs, err := repo.Branches()
	if err != nil {
		fmt.Println("Error fetching branches: " + err.Error())
		return nil, err
	}

	refs.ForEach(func(ref *plumbing.Reference) error {
		branches = append(branches, ref.Name().Short())
		return nil
	})

	return branches, nil
}

func getRemoteBranches(repoConf model.GitRepositoryConfig, repo *git.Repository) ([]string, error) {
	// ---- branches from the remote repository
	branches := make([]string, 0)
	remote, err := repo.Remote("origin")
	if err != nil {
		fmt.Println("Error fetching remote branches: " + err.Error())
		return nil, err
	}

	listOptions := &git.ListOptions{}
	publicKeys, err := getAuth(repoConf, err)
	if err != nil {
		return nil, err
	}
	if publicKeys != nil {
		listOptions.Auth = publicKeys
	}

	refs, err := remote.List(listOptions)
	if err != nil {
		fmt.Println("Error fetching remote branches: " + err.Error())
		return nil, err
	}

	for _, ref := range refs {
		branches = append(branches, ref.Name().Short())
	}

	// Sort the branches slice in increasing order.
	sort.Strings(branches)

	return branches, nil
}

// filterBranches filters the given branches and returns only those that are either "main" or start with "release-".
// It takes a slice of strings representing the branch names as input and returns a slice of strings containing the filtered branch names.
func filterBranches(branches []string) []string {
	releaseBranches := make([]string, 0)
	for _, branch := range branches {
		matched, _ := regexp.MatchString(`^(main|release-.*)$`, branch)
		if matched {
			releaseBranches = append(releaseBranches, branch)
		}
	}
	return releaseBranches
}
