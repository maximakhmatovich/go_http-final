package main

import (
	"log"
	"os"

	"github.com/maximakhmatovich/go_http-final/internal/server"
)

func main() {
	logsFile, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer logsFile.Close()

	logger := log.New(logsFile, "SERVER: ", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)

	newServer := server.NewServer(logger)

	logger.Printf("Запуск сервера на localhost: 8080")
	if err := newServer.ListenAndServe(); err != nil {
		logger.Fatal("Ошибка запуска сервера: ", err)
	}
}
