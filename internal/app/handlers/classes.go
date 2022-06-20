package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createClass(c *gin.Context) {
	id, _ := c.Get(userCtx)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
