package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/sport-point-backend.git/internal/app/config"
)

func (h *Handler) createClass(c *gin.Context) {
	id, ok := config.FromContextOk(c)
	if !ok {
		NewErrorResponse(c, http.StatusBadRequest, "failed to get ID")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
