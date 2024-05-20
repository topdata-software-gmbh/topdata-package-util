package util

import (
	"github.com/fatih/color"
	"log"
	"os"
)

// exists returns whether the given file or directory exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	log.Fatalln("Error checking if file or directory exists: " + err.Error())
	return false
}

func WriteToFile(file string, content string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

func RsyncDirectory(source string, destination string, exclude []string) {
	// ---- make sure both source and destination end with a slash
	if source[len(source)-1] != '/' {
		source += "/"
	}
	if destination[len(destination)-1] != '/' {
		destination += "/"
	}
	color.Yellow(">>>> RsyncDirectory: %s > %s ...", source, destination)

	// ---- build cmd
	cmd := []string{"rsync", "-avz" /*"--delete", */, source, destination}
	if len(exclude) > 0 {
		for i, _ := range exclude {
			cmd = append(cmd, "--exclude", exclude[i])
		}
	}
	// ---- action
	RunCommand(cmd[0], cmd[1:]...)
}
