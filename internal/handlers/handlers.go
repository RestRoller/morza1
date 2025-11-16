package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ServeMainPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func ProcessFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get uploaded file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return
	}

	content := string(contentBytes)
	if content == "" {
		http.Error(w, "Uploaded file is empty", http.StatusBadRequest)
		return
	}

	convertedContent, err := service.ConvertData(content)
	if err != nil {
		http.Error(w, "Conversion failed", http.StatusInternalServerError)
		return
	}

	fileExt := filepath.Ext(header.Filename)
	timestamp := time.Now().UTC().Format("20060102_150405")
	resultFilename := "converted_" + timestamp + fileExt

	resultFile, err := os.Create(resultFilename)
	if err != nil {
		http.Error(w, "Failed to create result file", http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()

	if _, err := resultFile.WriteString(convertedContent); err != nil {
		http.Error(w, "Failed to write result", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedContent))
}
