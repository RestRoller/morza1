package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type ApplicationServer struct {
	Log      *log.Logger
	HTTPServ *http.Server
}

func InitializeApplication(logger *log.Logger) *ApplicationServer {
	router := http.NewServeMux()

	// Main page route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not supported"))
			return
		}
		handlers.ProcessMainPage(w, r)
	})

	// File upload route
	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method not supported"))
			return
		}
		handlers.ProcessFileSubmission(w, r)
	})

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  8 * time.Second,
		WriteTimeout: 12 * time.Second,
		IdleTimeout:  20 * time.Second,
	}

	return &ApplicationServer{
		Log:      logger,
		HTTPServ: httpServer,
	}
}

func (as *ApplicationServer) StartServer() error {
	as.Log.Printf("Application server starting on %s", as.HTTPServ.Addr)
	return as.HTTPServ.ListenAndServe()
}
