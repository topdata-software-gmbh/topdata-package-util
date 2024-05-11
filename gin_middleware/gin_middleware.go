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

func PkgConfigListMiddleware(PkgConfigList *model.PkgConfigList) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("PkgConfigList", PkgConfigList)
		c.Next()
	}
}
