package handlers

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// IndexHandler обрабатывает запрос к корневому эндпоинту
func IndexHandler(w http.ResponseWriter, r *http.Request) {
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

// UploadHandler обрабатывает загрузку файлов
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Парсим форму
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error getting file from form: %v", err)
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	// Конвертируем содержимое
	converted, err := service.AutoDetectAndConvert(string(fileContent))
	if err != nil {
		log.Printf("Error converting content: %v", err)
		http.Error(w, "Unable to convert content", http.StatusInternalServerError)
		return
	}

	// Создаем локальный файл для результата
	originalExt := filepath.Ext(header.Filename)
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	outputFilename := "converted_" + timestamp + originalExt

	outputFile, err := os.Create(outputFilename)
	if err != nil {
		log.Printf("Error creating output file: %v", err)
		http.Error(w, "Unable to create output file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// Записываем результат в файл
	_, err = outputFile.WriteString(converted)
	if err != nil {
		log.Printf("Error writing to output file: %v", err)
		http.Error(w, "Unable to write to output file", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат клиенту
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	tmpl := template.Must(template.New("result").Parse(`
		<h2>Конвертация завершена!</h2>
		<p><strong>Исходный файл:</strong> {{.OriginalName}}</p>
		<p><strong>Результат сохранен в:</strong> {{.OutputName}}</p>
		<h3>Результат:</h3>
		<pre>{{.Content}}</pre>
		<a href="/">Назад</a>
	`))

	data := struct {
		OriginalName string
		OutputName   string
		Content      string
	}{
		OriginalName: header.Filename,
		OutputName:   outputFilename,
		Content:      converted,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Unable to display result", http.StatusInternalServerError)
		return
	}
}
