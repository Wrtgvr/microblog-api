package handlers

import (
	"github.com/wrtgvr/errsuit/drivers/ginadap"
)

type HandlerDeps struct {
	ErrHandler *ginadap.GinErrorHandler
}

func NewHandlerDeps(errHandler *ginadap.GinErrorHandler) *HandlerDeps {
	return &HandlerDeps{
		ErrHandler: errHandler,
	}
}
