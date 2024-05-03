package git_service_v2

// This service uses the git CLI to interact with git repositories.
// It is a replacement for the go-git library.

import (
	"fmt"
	"github.com/fatih/color"
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
