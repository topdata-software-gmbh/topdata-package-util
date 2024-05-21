package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRepositoriesHandler(c *gin.Context) {
	//PkgConfigList := c.MustGet("PkgConfigList").(*model.PkgConfigList)
	//
	//repoInfos, err := git_repo__old.GetRepoInfos(PkgConfigList.PkgConfigs, 10)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, "TODO: implement GetRepositoriesHandler")
}

func GetRepositoryDetailsHandler(c *gin.Context) {
	//PkgConfigList := c.MustGet("PkgConfigList").(*model.PkgConfigList)
	//
	//repoName := c.Param("name")
	//pkgConfig, err := git_repo__old.GetRepoDetails_old(repoName, PkgConfigList.PkgConfigs)
	//
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"error": err.Error(),
	//	})
	//	return
	//}
	//
	c.JSON(http.StatusOK, "TODO: implement GetRepositoryDetailsHandler")
}
