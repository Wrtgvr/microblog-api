package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/wrtgvr/errsuit/drivers/ginadap"
)

func MakeErrHandlerFunc(c *gin.Context) func(err error) bool {
	return ginadap.ErrHandlerFuncFromContext(c)
}
