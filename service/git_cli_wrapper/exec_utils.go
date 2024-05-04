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

func execCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		color.Yellow("!!!!! cmd: " + cmd.String())
		color.Red("!!!!! out: " + strings.TrimSpace(string(output)))
		color.Yellow("!!!!! err: " + err.Error())
		log.Fatalf("Failed to execute command: %s", err)
	}

	return string(output)
}

func execGitCommand(pkgConfig model.PkgConfig, args ...string) string {
	repoDir := pkgConfig.GetLocalGitRepoDir()
	args = append([]string{"-C", repoDir}, args...)
	// set environment variable for ssh key GIT_SSH_COMMAND='ssh -i /path/to/key'
	//if pkgConfig.PathSshKey != nil {
	//	fmt.Println("KEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEYKEYEY")
	//	args = append([]string{"-c", fmt.Sprintf("core.sshCommand='ssh -i %s'", *pkgConfig.PathSshKey)}, args...)
	//	fmt.Println("args: ", args)
	//}

	cmd := exec.Command("git", args...)

	// set environment variable for ssh key GIT_SSH_COMMAND='ssh -i /path/to/key'
	if pkgConfig.PathSshKey != nil {
		cmd.Env = append(os.Environ(), fmt.Sprintf("GIT_SSH_COMMAND=ssh -i %s", *pkgConfig.PathSshKey))
	}

	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))

	if err != nil {
		color.Yellow("!!!!!  cmd: " + cmd.String())
		color.Yellow("!!!!! code: " + err.Error())
		color.Red("!!!!!  out: " + strings.TrimSpace(string(output)))
		log.Fatalf("Failed to execute git command: %s", err)
	}

	return string(output)
}

func execShellCommand(command string) string {
	return execCommand("sh", "-c", command)
}
