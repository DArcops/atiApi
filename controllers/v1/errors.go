package v1

import (
	"net/http"

	"github.com/darcops/atiApi/models"
)

var (
	Err = map[error]int{
		models.ErrNotFound:   http.StatusNotFound,
		models.ErrToCreate:   http.StatusInternalServerError,
		models.ErrDuplicated: http.StatusConflict,
	}
)
