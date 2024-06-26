package pkg_zipper

import (
	"github.com/fatih/color"
	"github.com/topdata-software-gmbh/topdata-package-util/util"
	"path/filepath"
)

// CreateZipArchive creates a zip archive from the given source directory
func CreateZipArchive(sourceDir string, destZipFile string) {
	color.Yellow(">>>> CreateZipArchive: %s > %s", sourceDir, destZipFile)

	parentDir := filepath.Dir(sourceDir)
	sourceDirShort := filepath.Base(sourceDir)
	cmd := util.RenderString("cd {parentDir} && zip -FSr {destZipFile} {sourceDirShort}", map[string]string{
		"parentDir":      parentDir,
		"destZipFile":    destZipFile,
		"sourceDirShort": sourceDirShort,
	})

	util.RunShellCommand(cmd, nil)
}
