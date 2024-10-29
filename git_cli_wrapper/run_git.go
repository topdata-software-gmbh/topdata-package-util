package git_cli_wrapper

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"log"
	"os"
	"os/exec"
	"strings"
)

// runGitCommandInClonedRepo remains unchanged - keeps the fatal error handling
func runGitCommandInClonedRepo(pkgConfig *model.PkgConfig, args ...string) string {
	args = append([]string{"-C", pkgConfig.GetLocalGitRepoDir()}, args...)

	cmd := exec.Command("git", args...)

	// set environment variable for ssh key GIT_SSH_COMMAND='/usr/bin/ssh -i /path/to/key'
	if pkgConfig.PathSshKey != "" {
		extraEnv := fmt.Sprintf("GIT_SSH_COMMAND=/usr/bin/ssh -i %s", pkgConfig.PathSshKey)
		cmd.Env = append(os.Environ(), extraEnv)
	}

	color.Yellow(">>>> cmd: " + cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red("!!!!! code: " + err.Error())
		color.Red("!!!!!  out: " + strings.TrimSpace(string(output)))
		log.Fatalf("Failed to execute git command: %s", err)
	}

	return string(output)
}

// tryRunGitCommandInClonedRepo 10/2024 added - non-fatal variant that returns error instead of calling log.Fatalf
func tryRunGitCommandInClonedRepo(pkgConfig *model.PkgConfig, args ...string) (string, error) {
	args = append([]string{"-C", pkgConfig.GetLocalGitRepoDir()}, args...)

	cmd := exec.Command("git", args...)

	if pkgConfig.PathSshKey != "" {
		extraEnv := fmt.Sprintf("GIT_SSH_COMMAND=/usr/bin/ssh -i %s", pkgConfig.PathSshKey)
		cmd.Env = append(os.Environ(), extraEnv)
	}

	color.Yellow(">>>> cmd: " + cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red("!!!!! code: " + err.Error())
		color.Red("!!!!!  out: " + strings.TrimSpace(string(output)))
		return string(output), fmt.Errorf("git command failed: %w\nOutput: %s", err, string(output))
	}

	return string(output), nil
}
