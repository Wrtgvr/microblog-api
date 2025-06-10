package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/wrtgvr/microblog/internal/app"
	"github.com/wrtgvr/microblog/internal/config"
)

func main() {
	godotenv.Load()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := app.NewApp(cfg)

	go app.Server.MustRun()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	if err := app.Server.Stop(); err != nil {
		log.Printf("Shutdown error: %v\n", err)
	}
	// close db
	log.Println("Server stopped")
}
