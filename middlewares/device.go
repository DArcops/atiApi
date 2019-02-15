package middlewares

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/darcops/atiApi/models"
	"github.com/gin-gonic/gin"
)

func Device(c *gin.Context) {
	strID := c.Param("device_id")

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid device id"))
		return
	}

	device := &models.Device{ID: uint(id)}
	device.Get()

	c.Set("device", device)
	c.Next()
}
