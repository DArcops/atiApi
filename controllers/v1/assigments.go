package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/darcops/atiApi/models"
	"github.com/gin-gonic/gin"
)

func GetAssigments(c *gin.Context) {
	device := c.MustGet("device").(*models.Device)

	assigments, err := models.GetAssigments(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, assigments)
}

func AddAssigment(c *gin.Context) {
	var assigment = new(models.Assigment)
	device := c.MustGet("device").(*models.Device)

	if err := c.Bind(assigment); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := assigment.Create(device); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, assigment)
}

func GetAssigment(c *gin.Context) {
	strID := c.Param("assigment_id")

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid assigment id"))
		return
	}

	assigment := &models.Assigment{ID: uint(id)}

	if err := assigment.Get(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, assigment)
}

func DeleteAssigment(c *gin.Context) {
	device := c.MustGet("device").(*models.Device)
	strID := c.Param("assigment_id")

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.New("Invalid assigment id"))
		return
	}

	assigment := &models.Assigment{ID: uint(id)}

	if err := assigment.Delete(device); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
