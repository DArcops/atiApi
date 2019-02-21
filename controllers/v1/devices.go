package v1

import (
	"errors"
	"net/http"

	"github.com/darcops/atiApi/models"
	"github.com/gin-gonic/gin"
)

func GetDevices(c *gin.Context) {
	provider := c.MustGet("provider").(*models.Provider)

	devices, err := provider.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, devices)
}

func AddDevice(c *gin.Context) {
	var device = new(models.Device)

	if err := c.Bind(device); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := device.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusCreated, device)
}

func UploadDevicesFromFile(c *gin.Context) {
	provider := c.MustGet("provider").(*models.Provider)

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New(err.Error()))
		return
	}

	if err := models.SaveDevicesFromFile(file, provider); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, nil)
}
