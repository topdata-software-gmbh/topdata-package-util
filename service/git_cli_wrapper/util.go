package git_cli_wrapper

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"os/exec"
	"strings"
)

func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))

	if err != nil {
		color.Yellow("!!!!! cmd: " + cmd.String())
		color.Red("!!!!! out: " + strings.TrimSpace(string(output)))
		color.Yellow("!!!!! err: " + err.Error())

		return err
	}
	return nil
}

func execGitCommand(repoConfig model.PkgConfig, args ...string) (string, error) {
	repoDir := repoConfig.GetLocalGitRepoDir()
	args = append([]string{"-C", repoDir}, args...)
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))

	if err != nil {
		color.Yellow("!!!!!  cmd: " + cmd.String())
		color.Yellow("!!!!! code: " + err.Error())
		color.Red("!!!!!  out: " + strings.TrimSpace(string(output)))

		return string(output), err
	}

	return string(output), nil
}
