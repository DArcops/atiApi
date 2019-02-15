package routes

import (
	V1 "github.com/darcops/dialgorithm-server/controllers/v1"
	"github.com/darcops/dialgorithm-server/middlewares"
	"github.com/gin-gonic/gin"
)

var providers *gin.RouterGroup
var provider *gin.RouterGroup

func providerRoutes() {
	providers = v1.Group("providers")
	providers.GET("", V1.GetProviders)
	providers.POST("", V1.AddProvider)

	provider = providers.Group("/:provider_id")
	provider.Use(middlewares.Provider)
}
