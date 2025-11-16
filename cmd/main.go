package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "MORSE_CONVERTER: ", log.LstdFlags|log.Lshortfile)

	// Создаем сервер
	srv := server.New(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
