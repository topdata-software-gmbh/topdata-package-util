// In service/repositories.go
package git_repository_service

import (
	"github.com/topdata-software-gmbh/topdata-package-service/struct"
)

func GetRepositories(config _struct.Config) []_struct.GitRepository {
	return config.Repositories
}
