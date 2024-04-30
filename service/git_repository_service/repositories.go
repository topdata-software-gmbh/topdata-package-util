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
	"sync"
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

	return branches, nil
}

func refreshRepo(repoConf model.GitRepositoryConfig, destGitDir string) (*git.Repository, error) {

	var err error

	// Create a unique directory for each repoConf
	if err = os.MkdirAll(destGitDir, 0755); err != nil {
		return nil, err
	}

	publicKeys, err := getAuth(repoConf, err)
	if err != nil {
		return nil, err
	}

	var repo *git.Repository

	// Check if the repoConf has already been cloned
	if _, err = os.Stat(filepath.Join(destGitDir, ".git")); os.IsNotExist(err) {
		// If not, clone the repoConf
		fmt.Println(">>>> Cloning repoConf: " + repoConf.URL)

		cloneOptions := &git.CloneOptions{
			URL: repoConf.URL,
		}
		if publicKeys != nil {
			cloneOptions.Auth = publicKeys
		}

		repo, err = git.PlainClone(destGitDir, false, cloneOptions)
	} else {
		// If it has, open the existing repoConf
		fmt.Println(">>>> Using existing repoConf: " + destGitDir)
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

func getAuth(repoConf model.GitRepositoryConfig, err error) (*ssh.PublicKeys, error) {
	var publicKeys *ssh.PublicKeys = nil
	if repoConf.PathSshKey == nil {
		return nil, nil
	}
	fmt.Println(">>>> Using ssh key: " + *repoConf.PathSshKey)
	publicKeys, err = ssh.NewPublicKeysFromFile("git", *repoConf.PathSshKey, "")
	if err != nil {
		fmt.Println("Error reading ssh key: " + err.Error())
		return nil, err
	}
	// dump publicKeys
	// fmt.Println("loaded publicKeys: " + publicKeys.User + " " + publicKeys.Name())

	return publicKeys, nil
}

//func GetRepoInfos(repoConfigs []model.GitRepositoryConfig) ([]model.GitRepositoryInfo, error) {
//	repoInfos := make([]model.GitRepositoryInfo, len(repoConfigs))
//	// ---- fetch branches from the repoConfig
//	for i, repoConfig := range repoConfigs {
//		branches, err := GetRepositoryBranches(repoConfig)
//		if err != nil {
//			return nil, err
//		}
//		repoInfos[i].Name = repoConfig.Name
//		repoInfos[i].URL = repoConfig.URL
//		repoInfos[i].Description = repoConfig.Description
//		repoInfos[i].Branches = branches
//		//		repoInfos[i].ReleaseBranches =
//	}
//	return repoInfos, nil
//}

func GetRepoInfos(repoConfigs []model.GitRepositoryConfig) ([]model.GitRepositoryInfo, error) {
	var wg sync.WaitGroup
	repoInfoCh := make(chan model.GitRepositoryInfo, len(repoConfigs))
	errCh := make(chan error, len(repoConfigs))

	for _, repoConfig := range repoConfigs {
		wg.Add(1)
		go func(rc model.GitRepositoryConfig) {
			defer wg.Done()
			branches, err := GetRepositoryBranches(rc)
			if err != nil {
				errCh <- err
				return
			}
			repoInfoCh <- model.GitRepositoryInfo{
				Name:        rc.Name,
				URL:         rc.URL,
				Description: rc.Description,
				Branches:    branches,
			}
		}(repoConfig)
	}

	go func() {
		wg.Wait()
		close(repoInfoCh)
		close(errCh)
	}()

	repoInfos := make([]model.GitRepositoryInfo, 0, len(repoConfigs))
	for info := range repoInfoCh {
		repoInfos = append(repoInfos, info)
	}

	if len(errCh) > 0 {
		return nil, <-errCh
	}

	return repoInfos, nil
}
