package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags|log.Lshortfile)
	s := server.New(logger)

	err := s.HTTP.ListenAndServe()
	if err != nil {
		logger.Fatalf("Server error: %v", err)
	}
}
