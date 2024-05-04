package gin_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/topdata-software-gmbh/topdata-package-service/model"
)

func WebserverConfigMiddleware(webserverConfig model.WebserverConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("webserverConfig", webserverConfig)
		c.Next()
	}
}

func PkgConfigsMiddleware(pkgConfigs []model.PkgConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pkgConfigs", pkgConfigs)
		c.Next()
	}
}
