package pkg_zipper

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
)

// CreateZipArchive creates a zip archive from the given source directory
func CreateZipArchive(sourceDir string, destZipFile string) {
	color.Yellow(">>>> CreateZipArchive: %s > %s", sourceDir, destZipFile)
	util.RunCommand("zip", "-r", destZipFile, sourceDir, "-x", ".git")
}
