package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type HTTPServer struct {
	logger *log.Logger
	server *http.Server
}

func Initialize(logger *log.Logger) *HTTPServer {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.ServeMainPage)
	router.HandleFunc("/upload", handlers.ProcessFileUpload)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &HTTPServer{
		logger: logger,
		server: httpServer,
	}
}

func (s *HTTPServer) Start() error {
	s.logger.Printf("Starting HTTP server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}
