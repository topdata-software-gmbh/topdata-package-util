package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
	"github.com/topdata-software-gmbh/topdata-package-service/service/git_repo__old"
	"net/http"
)

func GetRepositoriesHandler(c *gin.Context) {
	serviceConfig := c.MustGet("serviceConfig").(model.ServiceConfig)

	repoInfos, err := git_repo__old.GetRepoInfos(serviceConfig.RepositoryConfigs, 10)
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
	repoConfig, err := git_repo__old.GetRepoDetails(repoName, serviceConfig.RepositoryConfigs)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, repoConfig)
}
