package middlewares

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/gin-gonic/gin"
)

func Provider(c *gin.Context) {
	strID := c.Param("provider_id")

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("invalid provider id"))
		return
	}

	provider := &models.Provider{ID: uint(id)}
	provider.Get()

	c.Set("provider", provider)
	c.Next()
}
