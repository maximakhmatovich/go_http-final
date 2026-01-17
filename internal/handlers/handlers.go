package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/maximakhmatovich/go_http-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Страница не найдена", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}
	text := string(content)

	result, err := service.Checker(text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newFileName := fmt.Sprintf("%s%s", time.Now().UTC().String(), filepath.Ext("myFile"))
	localFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(w, "Ошибка при создании файла", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	localFile.WriteString(fmt.Sprintf("%s\n", result))

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

	// "Вернуть результат конвертации строки." возможно придется вернуть файлом
}
