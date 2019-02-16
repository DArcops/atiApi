package routes

import (
	V1 "github.com/darcops/atiApi/controllers/v1"
	"github.com/gin-gonic/gin"
)

var assigments *gin.RouterGroup
var assigment *gin.RouterGroup

func assigmentsRoutes() {
	assigments = device.Group("assigments")

	assigments.GET("", V1.GetAssigments)
	assigments.POST("", V1.AddAssignment)

	assigment = assigments.Group("/:assigment_id")
	assigment.DELETE("", V1.DeleteAssigment)
	assigment.GET("", V1.GetAssigment)
}
