package git_repository_service

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os"
	"path/filepath"
	"sync"
)

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

func getAuth(repoConf model.GitRepositoryConfig, err error) (*ssh.PublicKeys, error) {
	var publicKeys *ssh.PublicKeys = nil
	if repoConf.PathSshKey == nil {
		return nil, nil
	}
	//fmt.Println(">>>> Using ssh key: " + *repoConf.PathSshKey)
	publicKeys, err = ssh.NewPublicKeysFromFile("git", *repoConf.PathSshKey, "")
	if err != nil {
		fmt.Println("!!!! Error reading ssh key: " + err.Error())
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

//func GetRepoInfos(repoConfigs []model.GitRepositoryConfig) ([]model.GitRepositoryInfo, error) {
//	var wg sync.WaitGroup
//	repoInfoCh := make(chan model.GitRepositoryInfo, len(repoConfigs))
//	errCh := make(chan error, len(repoConfigs))
//
//	for _, repoConfig := range repoConfigs {
//		wg.Add(1)
//		go func(rc model.GitRepositoryConfig) {
//			defer wg.Done()
//			branches, err := GetRepositoryBranches(rc)
//			if err != nil {
//				errCh <- err
//				return
//			}
//			repoInfoCh <- model.GitRepositoryInfo{
//				Name:        rc.Name,
//				URL:         rc.URL,
//				Description: rc.Description,
//				Branches:    branches,
//			}
//		}(repoConfig)
//	}
//
//	go func() {
//		wg.Wait()
//		close(repoInfoCh)
//		close(errCh)
//	}()
//
//	repoInfos := make([]model.GitRepositoryInfo, 0, len(repoConfigs))
//	for info := range repoInfoCh {
//		repoInfos = append(repoInfos, info)
//	}
//
//	if len(errCh) > 0 {
//		return nil, <-errCh
//	}
//
//	return repoInfos, nil
//}

// In Go, you can limit the number of parallel goroutines by using a buffered channel as a semaphore. Here's how you can do it:
// Create a buffered channel with a capacity equal to the maximum number of goroutines you want to allow to run concurrently.
// Before starting a new goroutine, send a value into the channel. This operation will block if the channel is already full, effectively limiting the number of concurrently running goroutines.
// When a goroutine finishes, read a value from the channel to allow another goroutine to start.
func GetRepoInfos(repoConfigs []model.GitRepositoryConfig, maxConcurrency int) ([]model.GitRepositoryInfo, error) {
	var wg sync.WaitGroup
	repoInfoCh := make(chan model.GitRepositoryInfo, len(repoConfigs))
	errCh := make(chan error, len(repoConfigs))

	// Create a buffered channel as a semaphore
	sem := make(chan struct{}, maxConcurrency)

	for _, repoConfig := range repoConfigs {
		wg.Add(1)

		// Send a value into the semaphore; this will block if the semaphore is full
		sem <- struct{}{}

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

			// Read a value from the semaphore, allowing another goroutine to proceed
			<-sem
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

func GetRepoDetails(repoName string, repoConfigs []model.GitRepositoryConfig) (model.GitRepositoryInfo, error) {
	for _, repoConfig := range repoConfigs {
		if repoConfig.Name == repoName {
			branches, err := GetRepositoryBranches(repoConfig)
			if err != nil {
				return model.GitRepositoryInfo{}, err
			}
			return model.GitRepositoryInfo{
				Name:            repoConfig.Name,
				URL:             repoConfig.URL,
				Branches:        branches,
				ReleaseBranches: filterBranches(branches),
			}, nil
		}
	}
	return model.GitRepositoryInfo{}, fmt.Errorf("repository not found: %s", repoName)
}
