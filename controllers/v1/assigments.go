package v1

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/darcops/atiApi/models"
	"github.com/gin-gonic/gin"
)

func GetAssigments(c *gin.Context) {
	p := c.MustGet("provider").(*models.Provider)

	assigments, err := models.GetAssigments(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, assigments)
}

func AddAssigment(c *gin.Context) {
	devices := []models.Device{}
	p := c.MustGet("provider").(*models.Provider)

	req := struct {
		Imeis        []string `json:"imeis" binding:"required"`
		AssignedUser string   `json:"user_name" binding:"required"`
		EndDate      string   `json:"end_date" binding:"required"`
		Description  string   `json:"description" binding:"required"`
	}{}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	for _, i := range req.Imeis {
		device := models.Device{}
		if models.First(&device, "imei = ? and provider_id = ?", i, p.ID).RecordNotFound() {
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(http.StatusNotFound, "imei not found for this provider")
			return
		}
		devices = append(devices, device)
	}

	var user = new(models.User)
	if models.First(user, "name = ?", req.AssignedUser).RecordNotFound() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusNotFound, "assigned user not found")
		return
	}

	assigment := &models.Assigment{
		Devices:     devices,
		Username:    req.AssignedUser,
		Description: req.Description,
		EndDate:     req.EndDate,
	}

	if err := assigment.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(http.StatusOK, assigment)
}

func GetAssigment(c *gin.Context) {
	strID := c.Param("assigment_id")

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusBadRequest, errors.New("Invalid assigment id"))
		return
	}

	assigment := &models.Assigment{ID: uint(id)}

	if err := assigment.Get(); err != nil {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
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
