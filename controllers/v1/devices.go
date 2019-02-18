package v1

import (
	"net/http"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/entropyx/sara/models/errors"
	"github.com/gin-gonic/gin"
)

func GetDevices(c *gin.Context) {
	provider := c.MustGet("provider").(*models.Provider)

	devices, err := provider.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

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
	c.JSON(http.StatusCreated, device)
}

func UploadDevicesFromFile(c *gin.Context) {
	provider := c.MustGet("provider").(*models.Provider)

	file, err := c.FormFile("file")
	if err != nil {
		errors.JSON(c, errors.BadRequest(err.Error()))
		return
	}

	if err := models.SaveDevicesFromFile(file, provider); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
