package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/maximakhmatovich/go_http-final/internal/service"
)

var logger *log.Logger

func SetLogger(l *log.Logger) {
	logger = l
	if logger == nil {
		logger = log.Default()
	}
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		logger.Println(err)
		http.Error(w, "Страница не найдена", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(w, file)
	if err != nil {
		logger.Println(err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Println("Метод не разрешен:", r.Method)
		http.Error(w, "Метод не разрешен", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		logger.Println(err)
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		logger.Println(err)
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	text := string(content)

	result, err := service.Checker(text)
	if err != nil {
		logger.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newFileName := filepath.Join("uploads", fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(header.Filename)))

	if err := os.MkdirAll("uploads", 0755); err != nil {
		logger.Println(err)
		http.Error(w, "Ошибка при создании папки uploads", http.StatusInternalServerError)
		return
	}

	localFile, err := os.Create(newFileName)
	if err != nil {
		logger.Println(err)
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	if _, err := localFile.WriteString(fmt.Sprintf("%s\n", result)); err != nil {
		logger.Println(err)
		http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

	// "Вернуть результат конвертации строки." возможно придется вернуть файлом
}
