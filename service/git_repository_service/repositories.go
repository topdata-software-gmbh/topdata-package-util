// In service/repositories.go
package git_repository_service

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
	"path/filepath"
)

func GetRepositories(config model.Config) []model.GitRepository {
	return config.Repositories
}

func GetRepositoryBranches(repoModel model.GitRepository) ([]string, error) {
	fmt.Println("GetRepositoryBranches: " + repoModel.Name)
	// ---- fetch branches from the repoModel
	gitDir := filepath.Join("/tmp/git-repos", repoModel.Name)

	// ---- git clone / pull
	repo, err := refreshRepo(repoModel, gitDir)
	if err != nil {
		fmt.Println("Error fetching repo to " + gitDir + ": " + err.Error())
		return nil, err
	}

	//// ---- fetch branches from the repoModel
	//branches, err := getBranches(repo)
	//if err != nil {
	//	fmt.Println("Error fetching branches to " + gitDir + ": " + err.Error())
	//	return nil
	//}
	//return branches

	remoteBranches, err := getRemoteBranches(repoModel, repo)
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

func getRemoteBranches(repoModel model.GitRepository, repo *git.Repository) ([]string, error) {
	// ---- branches from the remote repository
	branches := make([]string, 0)
	remote, err := repo.Remote("origin")
	if err != nil {
		fmt.Println("Error fetching remote branches: " + err.Error())
		return nil, err
	}

	listOptions := &git.ListOptions{}
	publicKeys, err := getAuth(repoModel, err)
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

	return branches, nil
}

func refreshRepo(repository model.GitRepository, destGitDir string) (*git.Repository, error) {

	var err error

	// Create a unique directory for each repository
	if err = os.MkdirAll(destGitDir, 0755); err != nil {
		return nil, err
	}

	publicKeys, err := getAuth(repository, err)
	if err != nil {
		return nil, err
	}

	var repo *git.Repository

	// Check if the repository has already been cloned
	if _, err = os.Stat(filepath.Join(destGitDir, ".git")); os.IsNotExist(err) {
		// If not, clone the repository
		fmt.Println(">>>> Cloning repository: " + repository.URL)

		cloneOptions := &git.CloneOptions{
			URL: repository.URL,
		}
		if publicKeys != nil {
			cloneOptions.Auth = publicKeys
		}

		repo, err = git.PlainClone(destGitDir, false, cloneOptions)
	} else {
		// If it has, open the existing repository
		fmt.Println(">>>> Using existing repository: " + destGitDir)
		repo, err = git.PlainOpen(destGitDir)
		if err == nil {
			// And pull the latest changes from the origin remote
			worktree, err := repo.Worktree()
			if err != nil {
				return nil, err
			}

			pullOptions := &git.PullOptions{
				RemoteName: "origin",
			}
			if publicKeys != nil {
				pullOptions.Auth = publicKeys
			}

			err = worktree.Pull(pullOptions)
			if err != nil && err != git.NoErrAlreadyUpToDate {
				return nil, err
			}
		}
	}

	return repo, err
}

func getAuth(repository model.GitRepository, err error) (*ssh.PublicKeys, error) {
	var publicKeys *ssh.PublicKeys = nil
	if repository.PathSshKey == nil {
		return nil, nil
	}
	fmt.Println(">>>> Using ssh key: " + *repository.PathSshKey)
	publicKeys, err = ssh.NewPublicKeysFromFile("git", *repository.PathSshKey, "")
	if err != nil {
		fmt.Println("Error reading ssh key: " + err.Error())
		return nil, err
	}
	// dump publicKeys
	// fmt.Println("loaded publicKeys: " + publicKeys.User + " " + publicKeys.Name())

	return publicKeys, nil
}
