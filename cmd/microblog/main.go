package main

import (
	"github.com/wrtgvr/microblog/internal/app"
)

func main() {
	app := app.NewApp(8080)
	app.Server.MustRun()

	// graceful shutdown
}
