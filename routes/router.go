package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

var (
	router *gin.Engine
	api    *gin.RouterGroup
	v1     *gin.RouterGroup
)

func init() {
	router = gin.New()
	api = router.Group("api")
	v1 = api.Group("v1")

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,POST,DELETE,PUT",
		RequestHeaders:  "Origin, Authorization, Content-Type, Access-Control-Allow-Origin",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	userRoutes()
	providerRoutes()
	devicesRoutes()
	assigmentsRoutes()

	router.Run(":8088")
}
