package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func ProcessMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func ProcessFileSubmission(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	parseErr := r.ParseMultipartForm(32 << 20)
	if parseErr != nil {
		errorMsg := fmt.Sprintf("Form parsing error: %v", parseErr)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}

	uploadedFile, fileHeader, fileErr := r.FormFile("uploadedFile")
	if fileErr != nil {
		errorMsg := fmt.Sprintf("File access error: %v", fileErr)
		http.Error(w, errorMsg, http.StatusBadRequest)
		return
	}
	defer func() {
		if closeErr := uploadedFile.Close(); closeErr != nil {
			fmt.Printf("File close error: %v\n", closeErr)
		}
	}()

	var contentBuf bytes.Buffer
	_, copyErr := io.Copy(&contentBuf, uploadedFile)
	if copyErr != nil {
		errorMsg := fmt.Sprintf("Content reading error: %v", copyErr)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	fileData := contentBuf.String()
	if len(fileData) == 0 {
		http.Error(w, "Empty file content", http.StatusBadRequest)
		return
	}

	processedResult := service.ConvertContent(fileData)

	timestamp := time.Now().UTC().Format("20060102_150405")
	originalFilename := fileHeader.Filename
	outputFilename := "result_" + timestamp + "_" + originalFilename

	writeErr := os.WriteFile(outputFilename, []byte(processedResult), 0600)
	if writeErr != nil {
		errorMsg := fmt.Sprintf("File creation error: %v", writeErr)
		http.Error(w, errorMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(processedResult))
}
