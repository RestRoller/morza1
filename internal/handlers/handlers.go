package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			// Логируем, но не прерываем выполнение
		}
	}()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	content := string(fileContent)
	if content == "" {
		http.Error(w, "File is empty", http.StatusBadRequest)
		return
	}

	converted := service.ToggleMorse(content)

	originalExt := filepath.Ext(header.Filename)
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	outputFilename := "converted_" + timestamp + originalExt

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		http.Error(w, "Unable to create output file", http.StatusInternalServerError)
		return
	}
	defer func() {
		if closeErr := outputFile.Close(); closeErr != nil {
			// Логируем, но не прерываем выполнение
		}
	}()

	_, err = outputFile.WriteString(converted)
	if err != nil {
		http.Error(w, "Unable to write to output file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(converted))
	if err != nil {
		// Игнорируем ошибку записи в response, так как клиент мог разорвать соединение
		return
	}
}
