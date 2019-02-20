package v1

import (
	"net/http"

	"github.com/darcops/atiApi/models"
	"github.com/gin-gonic/gin"
)

func GetProviders(c *gin.Context) {
	providers, err := models.GetProviders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, providers)
}

func GetProvider(c *gin.Context) {
	provider := c.MustGet("provider").(*models.Provider)

	c.JSON(http.StatusOK, provider)
}

func AddProvider(c *gin.Context) {
	var provider = new(models.Provider)

	if err := c.Bind(provider); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := provider.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, provider)
}
