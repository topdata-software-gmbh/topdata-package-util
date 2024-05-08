package git_repo__old

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"regexp"
	"sort"
)

func GetRepositoryBranches_old(repoConfig model.PkgConfig) ([]string, error) {
	fmt.Println(">>>> GetRepositoryBranches_old: " + repoConfig.Name)
	// ---- fetch branches from the repoConfig
	gitDir := repoConfig.GetLocalGitRepoDir()

	// ---- git clone / pull
	repo, err := refreshRepo_old(repoConfig, gitDir)
	if err != nil {
		fmt.Println("Error fetching repo to " + gitDir + ": " + err.Error())
		return nil, err
	}

	remoteBranches, err := getRemoteBranches(repoConfig, repo)
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

func getRemoteBranches(repoConf model.PkgConfig, repo *git.Repository) ([]string, error) {
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

// FilterBranches_old filters the given branches and returns only those that are either "server" or start with "release-".
// It takes a slice of strings representing the branch names as input and returns a slice of strings containing the filtered branch names.
func FilterBranches_old(branches []string, regexPattern string) []string {
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

func GetCommitId(repoConfig model.PkgConfig, branchName string) (string, error) {
	destGitDir := repoConfig.GetLocalGitRepoDir()
	repo, err := refreshRepo_old(repoConfig, destGitDir)

	if err != nil {
		log.Println("Error fetching repo to " + destGitDir + ": " + err.Error())
		return "", err
	}

	checkoutBranch(repo, branchName)

	// Get the commit ID for the branchName
	ref, err := repo.Reference(plumbing.NewBranchReferenceName(branchName), true)
	if err != nil {
		return "", err
	}
	return ref.Hash().String(), nil
}

func checkoutBranch(r *git.Repository, branchName string) {

	// ... checking out branchName
	log.Println("git checkout %s", branchName)

	worktree, err := r.Worktree()
	CheckIfError(err)

	branchRefName := plumbing.NewBranchReferenceName(branchName)
	branchCoOpts := git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branchRefName),
		Force:  true,
		Keep:   false,
		Create: true,
	}

	err = worktree.Checkout(&branchCoOpts)
	if err != nil {
		log.Println("local checkout of branchName '%s' failed, will attempt to fetch remote branchName of same name.", branchName)
		log.Println("like `git checkout <branchName>` defaulting to `git checkout -b <branchName> --track <remote>/<branchName>`")

		mirrorRemoteBranchRefSpec := fmt.Sprintf("refs/heads/%s:refs/heads/%s", branchName, branchName)
		err = fetchOrigin(r, mirrorRemoteBranchRefSpec)
		CheckIfError(err)

		err = worktree.Checkout(&branchCoOpts)
		CheckIfError(err)
	}
	CheckIfError(err)

	log.Println("checked out branchName: %s", branchName)

	// ... retrieving the commit being pointed by HEAD (branchName now)
	log.Println("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())
}

func fetchOrigin(repo *git.Repository, refSpecStr string) error {
	remote, err := repo.Remote("origin")
	CheckIfError(err)

	var refSpecs []config.RefSpec
	if refSpecStr != "" {
		refSpecs = []config.RefSpec{config.RefSpec(refSpecStr)}
	}

	if err = remote.Fetch(&git.FetchOptions{
		RefSpecs: refSpecs,
	}); err != nil {
		if err == git.NoErrAlreadyUpToDate {
			fmt.Print("refs already up to date")
		} else {
			return fmt.Errorf("fetch origin failed: %v", err)
		}
	}

	return nil
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatalf("An error occurred: %s", err)
	}
}
