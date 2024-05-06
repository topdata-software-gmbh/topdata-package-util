package model

import (
	"encoding/json"
	"github.com/fatih/color"
	"log"
	"os"
)

// PkgInfoList - just a container for a list of PkgInfo with convenience search functionality
type PkgInfoList struct {
	PkgInfos []PkgInfo
}

// FindOneByNameOrFail - find a package by name
func (rcl *PkgInfoList) FindOneByNameOrFail(name string) *PkgInfo {
	for _, pkgInfo := range rcl.PkgInfos {
		if pkgInfo.Name == name {
			return &pkgInfo
		}
	}
	log.Fatalln("Package not found: " + name)
	return nil
}

func (pil *PkgInfoList) SaveToDisk(filePath string) {
	color.Yellow(">>>> Saving to pkInfoList to %s", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file %s: %s\n", filePath, err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(pil)
	if err != nil {
		log.Fatalf("Error encoding to file %s: %s\n", filePath, err.Error())
	}
}

func LoadFromDisk(filePath string) (*PkgInfoList, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	pkgInfoList := &PkgInfoList{}
	err = decoder.Decode(pkgInfoList)
	if err != nil {
		return nil, err
	}

	return pkgInfoList, nil
}
