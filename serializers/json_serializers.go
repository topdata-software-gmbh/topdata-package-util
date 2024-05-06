package serializers

import (
	"encoding/json"
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"log"
	"os"
)

func SavePkgInfoList(pil *model.PkgInfoList, filePath string) {
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

func LoadPkgInfoList(filePath string) *model.PkgInfoList {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file %s: %s\n", filePath, err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	pkgInfoList := &model.PkgInfoList{}
	err = decoder.Decode(pkgInfoList)
	if err != nil {
		log.Fatalf("Error decoding file %s: %s\n", filePath, err.Error())
	}

	return pkgInfoList
}
