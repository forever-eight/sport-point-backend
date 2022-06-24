package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

type createClassResponse struct {
	*ds.ClassOutput
}

func (h *Handler) createClass(c *gin.Context) {
	/*	id, ok := config.FromContextOk(c)
		if !ok {
			NewErrorResponse(c, http.StatusBadRequest, "failed to get ID")
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"id": id,
		})*/
	var input ds.ClassInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	class, err := h.Service.CreateClass(&input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, createClassResponse{class})
}
