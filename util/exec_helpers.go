package util

import (
	"github.com/fatih/color"
	"log"
	"os/exec"
	"strings"
)

// RunCommand executes a (CLI) command
func RunCommand(command string, args ...string) string {
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

// RunShellCommand executes a shell. Chaining commands with pipes is possible. passing extraEnv is optional.
func RunShellCommand(shellCommand string, extraEnv []string) string {
	//color.Yellow("================= RunShellCommand env: %s", extraEnv)
	//color.Yellow("================= RunShellCommand cmd: %s", shellCommand)

	cmd := exec.Command("/usr/bin/sh", "-c", shellCommand)
	cmd.Env = extraEnv

	output, err := cmd.CombinedOutput()

	if err != nil {
		color.Yellow("!!!!! cmd: " + cmd.String())
		color.Red("!!!!! out: " + strings.TrimSpace(string(output)))
		color.Yellow("!!!!! err: " + err.Error())
		log.Fatalf("Failed to execute shellCommand: %s", err)
	}

	return string(output)
}
