package handlers

import (
	"github.com/gin-gonic/gin"
	errsuit "github.com/wrtgvr/errsuit/core"
)

func DecodeReqBody(c *gin.Context, obj any) *errsuit.AppError {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		return errsuit.NewBadRequest("invalid body", err, false)
	}
	return nil
}
