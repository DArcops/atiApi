package routes

import (
	V1 "github.com/darcops/atiApi/controllers/v1"
	"github.com/gin-gonic/gin"
)

var devices *gin.RouterGroup
var device *gin.RouterGroup

func devicesRoutes() {
	devices = provider.Group("devices")

	devices.GET("", V1.GetDevices)
	devices.POST("", V1.AddDevice)
	devices.POST("/upload_file", V1.UploadDevicesFromFile)

	device = devices.Group("/:device_id")
	device.Use(middlewares.Device)
}
