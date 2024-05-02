package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_repository_service"
	"net/http"
)

func GetRepositoriesHandler(c *gin.Context) {
	serviceConfig := c.MustGet("serviceConfig").(model.ServiceConfig)

	repoInfos, err := git_repository_service.GetRepoInfos(serviceConfig.RepositoryConfigs, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, repoInfos)
}

func GetRepositoryDetailsHandler(c *gin.Context) {
	serviceConfig := c.MustGet("serviceConfig").(model.ServiceConfig)

	repoName := c.Param("name")
	for _, repoConfig := range serviceConfig.RepositoryConfigs {
		if repoConfig.Name == repoName {
			c.JSON(http.StatusOK, repoConfig)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Repository not found",
	})
}
