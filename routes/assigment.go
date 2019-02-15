package routes

import (
	V1 "github.com/darcops/dialgorithm-server/controllers/v1"
	"github.com/gin-gonic/gin"
)

var assigments *gin.RouterGroup

func assigmentsRoutes() {
	assigments = device.Group("assigments")

	assigments.GET("", V1.GetAssigments)
	assigments.POST("", V1.AddAssignment)
	assigments.DELETE("", V1.DeleteAssigment)
}
