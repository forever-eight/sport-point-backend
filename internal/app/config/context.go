package config

import (
	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func FromContext(c *gin.Context) *uint32 {
	// get id interface from ctx
	id, exists := c.Get(userCtx)
	if !exists {
		//handlers.NewErrorResponse(c, http.StatusBadRequest, "get id interface from ctx error")
		return nil
	}

	// upcast interface to uint32
	idUint, ok := id.(uint32)
	if !ok {
		//handlers.NewErrorResponse(c, http.StatusBadRequest, "upcast interface to uint32 error")
		return nil
	}

	return &idUint
}

func FromContextOk(gCtx *gin.Context) (*uint32, bool) {
	userID := FromContext(gCtx)

	return userID, userID != nil
}

func WrapContext(ctx *gin.Context, userID uint32) {
	ctx.Set(userCtx, userID)
}
