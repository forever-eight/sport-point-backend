package mw

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/forever-eight/sport-point-backend.git/internal/app/config"
	"github.com/forever-eight/sport-point-backend.git/internal/app/service"
)

const (
	authorizationHeader = "Authorization"
)

func UserIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			//app.NewErrorResponse(c, http.StatusBadRequest, "incorrect jwt")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			//app.NewErrorResponse(c, http.StatusBadRequest, "incorrect jwt(2 parts)")
			return
		}

		userId, err := service.ParseToken(headerParts[1])
		if err != nil {
			//app.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
			return
		}

		config.WrapContext(c, userId)
		c.Next()
	}
}
