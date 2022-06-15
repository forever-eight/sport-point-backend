package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

type createUserResponse struct {
	Result *ds.UserOutput `json:"result,omitempty"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input *ds.UserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	output, err := h.Service.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createUserResponse{Result: output})
}

func (h *Handler) signIn(c *gin.Context) {

}
