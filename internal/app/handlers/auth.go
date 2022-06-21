package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/sport-point-backend.git/internal/app/ds"
)

type signUpResponse struct {
	*ds.UserOutput
}

func (h *Handler) signUp(c *gin.Context) {
	var input *ds.UserInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	output, err := h.Service.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, signUpResponse{output})
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string
}

func (h *Handler) signIn(c *gin.Context) {
	var input *signInInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.Service.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, signInResponse{Token: token})
}
