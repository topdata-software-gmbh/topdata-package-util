package util

import (
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
