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
