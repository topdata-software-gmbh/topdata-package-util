package git_repo__old

//
////
////import (
////	"fmt"
////	"github.com/go-git/go-git/v5"
////	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
////	"github.com/topdata-software-gmbh/topdata-package-util/model"
////	"log"
////	"os"
////	"path/filepath"
////	"sync"
////)
////
//func refreshRepo_old(pkgConfig *model.PkgConfig, destGitDir string) (*git.Repository, error) {
//
//	var err error
//
//	// Create a unique directory for each pkgConfig
//	if err = os.MkdirAll(destGitDir, 0755); err != nil {
//		return nil, err
//	}
//
//	publicKeys, err := getAuth(pkgConfig, err)
//	if err != nil {
//		return nil, err
//	}
//
//	var repo *git.Repository
//
//	// Check if the pkgConfig has already been cloned
//	if _, err = os.Stat(filepath.Join(destGitDir, ".pkg")); os.IsNotExist(err) {
//		// If not, clone the pkgConfig
//		fmt.Println(">>>> Cloning pkgConfig: " + pkgConfig.URL)
//
//		cloneOptions := &git.CloneOptions{
//			URL: pkgConfig.URL,
//		}
//		if publicKeys != nil {
//			cloneOptions.Auth = publicKeys
//		}
//
//		repo, err = git.PlainClone(destGitDir, false, cloneOptions)
//	} else {
//		// If it has, open the existing pkgConfig
//		fmt.Println(">>>> Using existing repository: " + destGitDir)
//		repo, err = git.PlainOpen(destGitDir)
//		if err == nil {
//
//			fetchOptions := &git.FetchOptions{
//				RemoteName: "origin",
//				Force:      true,
//			}
//			if publicKeys != nil {
//				fetchOptions.Auth = publicKeys
//			}
//
//			err = repo.Fetch(fetchOptions)
//			if err != nil && err != git.NoErrAlreadyUpToDate {
//				return nil, err
//			}
//			// And pull the latest changes from the origin remote
//			//worktree, err := repo.Worktree()
//			//if err != nil {
//			//	return nil, err
//			//}
//			//pullOptions := &pkg.PullOptions{
//			//	RemoteName: "origin",
//			//	Force:      true,
//			//}
//			//if publicKeys != nil {
//			//	pullOptions.Auth = publicKeys
//			//}
//			//
//			//err = worktree.Pull(pullOptions)
//			//if err != nil && err != pkg.NoErrAlreadyUpToDate {
//			//	return nil, err
//			//}
//		}
//	}
//
//	return repo, err
//}
//
//func getAuth(pkgConfig *model.PkgConfig, err error) (*ssh.PublicKeys, error) {
//	var publicKeys *ssh.PublicKeys = nil
//	if pkgConfig.PathSshKey == nil {
//		return nil, nil
//	}
//	//fmt.Println(">>>> Using ssh key: " + *pkgConfig.PathSshKey)
//	publicKeys, err = ssh.NewPublicKeysFromFile("pkg", *pkgConfig.PathSshKey, "")
//	if err != nil {
//		fmt.Println("!!!! Error reading ssh key: " + err.Error())
//		return nil, err
//	}
//	// dump publicKeys
//	// fmt.Println("loaded publicKeys: " + publicKeys.User + " " + publicKeys.MachineName())
//
//	return publicKeys, nil
//}
//
////func GetRepoInfos(pkgConfigs []model.PkgConfig) ([]model.PkgInfo, error) {
////	repoInfos := make([]model.PkgInfo, len(pkgConfigs))
////	// ---- fetch branches from the pkgConfig
////	for i, pkgConfig := range pkgConfigs {
////		branches, err := GetRepositoryBranches_old(pkgConfig)
////		if err != nil {
////			return nil, err
////		}
////		repoInfos[i].MachineName = pkgConfig.MachineName
////		repoInfos[i].URL = pkgConfig.URL
////		repoInfos[i].Description = pkgConfig.Description
////		repoInfos[i].BranchNames = branches
////		//		repoInfos[i].ReleaseBranches =
////	}
////	return repoInfos, nil
////}
//
////func GetRepoInfos(pkgConfigs []model.PkgConfig) ([]model.PkgInfo, error) {
////	var wg sync.WaitGroup
////	repoInfoCh := make(chan model.PkgInfo, len(pkgConfigs))
////	errCh := make(chan error, len(pkgConfigs))
////
////	for _, pkgConfig := range pkgConfigs {
////		wg.Add(1)
////		go func(rc model.PkgConfig) {
////			defer wg.Done()
////			branches, err := GetRepositoryBranches_old(rc)
////			if err != nil {
////				errCh <- err
////				return
////			}
////			repoInfoCh <- model.PkgInfo{
////				MachineName:        rc.MachineName,
////				URL:         rc.URL,
////				Description: rc.Description,
////				BranchNames:    branches,
////			}
////		}(pkgConfig)
////	}
////
////	go func() {
////		wg.Wait()
////		close(repoInfoCh)
////		close(errCh)
////	}()
////
////	repoInfos := make([]model.PkgInfo, 0, len(pkgConfigs))
////	for info := range repoInfoCh {
////		repoInfos = append(repoInfos, info)
////	}
////
////	if len(errCh) > 0 {
////		return nil, <-errCh
////	}
////
////	return repoInfos, nil
////}
//
//// In Go, you can limit the number of parallel goroutines by using a buffered channel as a semaphore. Here's how you can do it:
//// Create a buffered channel with a capacity equal to the maximum number of goroutines you want to allow to run concurrently.
//// Before starting a new goroutine, send a value into the channel. This operation will block if the channel is already full, effectively limiting the number of concurrently running goroutines.
//// When a goroutine finishes, read a value from the channel to allow another goroutine to start.
//func GetRepoInfos(pkgConfigs []model.PkgConfig, maxConcurrency int) ([]model.PkgInfo, error) {
//	var wg sync.WaitGroup
//	repoInfoCh := make(chan model.PkgInfo, len(pkgConfigs))
//	errCh := make(chan error, len(pkgConfigs))
//
//	// Create a buffered channel as a semaphore
//	sem := make(chan struct{}, maxConcurrency)
//
//	for _, pkgConfig := range pkgConfigs {
//		wg.Add(1)
//
//		// Send a value into the semaphore; this will block if the semaphore is full
//		sem <- struct{}{}
//
//		go func(pkgConfig model.PkgConfig) {
//			defer wg.Done()
//
//			//branches, err := GetRepositoryBranches_old(pkgConfig)
//			//if err != nil {
//			//	errCh <- err
//			//	return
//			//}
//
//			repoInfoCh <- model.PkgInfo{
//				// Name:        pkgConfig.Name,
//				//URL:         pkgConfig.URL,
//				//Description: pkgConfig.Description,
//				PkgConfig: &pkgConfig,
//				// BranchNames: branches,
//			}
//
//			// Read a value from the semaphore, allowing another goroutine to proceed
//			<-sem
//		}(pkgConfig)
//	}
//
//	go func() {
//		wg.Wait()
//		close(repoInfoCh)
//		close(errCh)
//	}()
//
//	repoInfos := make([]model.PkgInfo, 0, len(pkgConfigs))
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
//
//func GetRepoDetails_old(repoName string, pkgConfigs []model.PkgConfig) (model.PkgInfo, error) {
//	for _, pkgConfig := range pkgConfigs {
//		branches, err := GetRepositoryBranches_old(pkgConfig)
//		if err != nil {
//			return model.PkgInfo{}, err
//		}
//		releaseBranchNames := FilterBranches_old(branches, `^(server|server-.*|release-.*)$`)
//
//		log.Println("releaseBranchNames: ", releaseBranchNames)
//
//		releaseBranches := make([]model.GitBranchInfo, 0)
//		// iterate over release branches and get pkg commit id for each
//		for _, branch := range releaseBranchNames {
//			// get commit id for the branch
//			commitId, err := GetCommitId(pkgConfig, branch)
//			if err != nil {
//				log.Println("Error getting commit ID for branch: " + branch + " " + err.Error())
//				return model.PkgInfo{}, err
//			}
//			releaseBranches = append(releaseBranches, model.GitBranchInfo{
//				Name:     branch,
//				CommitId: commitId,
//			})
//		}
//		return model.PkgInfo{
//			PkgConfig: pkgConfig,
//			//Name: pkgConfig.Name,
//			//URL:  pkgConfig.URL,
//			// BranchNames:     branches,
//			// ReleaseBranches: releaseBranches,
//		}, nil
//	}
//	return model.PkgInfo{}, fmt.Errorf("repository not found: %s", repoName)
//}
