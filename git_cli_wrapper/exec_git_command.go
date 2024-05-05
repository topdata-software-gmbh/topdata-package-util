package git_cli_wrapper

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"os"
	"os/exec"
	"strings"
)

func execGitCommand(pkgConfig model.PkgConfig, args ...string) string {
	repoDir := pkgConfig.GetLocalGitRepoDir()
	args = append([]string{"-C", repoDir}, args...)

	cmd := exec.Command("git", args...)

	// set environment variable for ssh key GIT_SSH_COMMAND='/usr/bin/ssh -i /path/to/key'
	if pkgConfig.PathSshKey != nil {
		extraEnv := fmt.Sprintf("GIT_SSH_COMMAND=/usr/bin/ssh -i %s", *pkgConfig.PathSshKey)
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
