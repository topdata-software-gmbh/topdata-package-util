package git_cli_wrapper

import (
	"github.com/topdata-software-gmbh/topdata-package-util/model"
	"log"
)

// GetLastCommitToMainAt it returns the date of the last commit to the `main` branch
func GetLastCommitToMainAt(pkgConfig *model.PkgConfig) string {
	err := TryCheckoutBranch(pkgConfig, "main")
	if err != nil {
		log.Printf("Failed to checkout branch: %v", err)
		return ""
	}

	return GetCommitDate(pkgConfig)
}
