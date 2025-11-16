package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	appLogger := log.New(os.Stdout, "APP: ", log.LstdFlags|log.Lshortfile)

	webServer := server.InitializeApplication(appLogger)

	startErr := webServer.StartServer()
	if startErr != nil {
		appLogger.Fatalf("Server startup error: %v", startErr)
	}
}
