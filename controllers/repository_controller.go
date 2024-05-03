package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/trash____git_repository"
	"net/http"
)

func GetRepositoriesHandler(c *gin.Context) {
	serviceConfig := c.MustGet("serviceConfig").(model.ServiceConfig)

	repoInfos, err := trash____git_repository.GetRepoInfos(serviceConfig.RepositoryConfigs, 10)
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
	repoConfig, err := trash____git_repository.GetRepoDetails(repoName, serviceConfig.RepositoryConfigs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, repoConfig)
}
