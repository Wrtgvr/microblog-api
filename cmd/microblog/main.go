package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wrtgvr/microblog/internal/app"
)

func main() {
	app := app.NewApp(8080)

	go app.Server.MustRun()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	if err := app.Server.Stop(); err != nil {
		log.Printf("Shutdown error: %v\n", err)
	}
	log.Println("Server stopped")
}
