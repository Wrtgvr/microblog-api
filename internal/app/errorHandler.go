package app

import (
	"fmt"
	"os"

	errsuit "github.com/wrtgvr/errsuit/core"
	"github.com/wrtgvr/errsuit/drivers/ginadap"
)

func prepareGinErrorHandler() *ginadap.GinErrorHandler {
	l := prepareErrorLogger()
	cfg := errsuit.Config{
		Format: errsuit.ResponseFormatJSON,
		Logger: l,
	}
	return ginadap.NewGinErrorHandler(cfg)
}

func prepareErrorLogger() *errsuit.Logger {
	return errsuit.NewLogger(os.Stdout, func(err error) string {
		appErr := errsuit.AsAppError(err)
		return fmt.Sprintf("ERROR: %s\n%v", appErr.Err.Error(), appErr.Err)
	})
}
