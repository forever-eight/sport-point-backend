package handlers

import (
	"github.com/gin-gonic/gin"
)

type errorType struct {
	Message string `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorType{Message: message})
}
