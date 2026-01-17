package service

import (
	"errors"
	"log"
	"strings"

	"github.com/maximakhmatovich/go_http-final/pkg/morse"
)

var (
	ErrInvalidInput = errors.New("Пустая строка на входе")
	logger          *log.Logger
)

func SetLogger(l *log.Logger) {
	logger = l
	if logger == nil {
		logger = log.Default()
	}
}

func isMorse(text string) (bool, error) {
	if text == "" {
		logger.Println("Пустая строка на входе")
		return false, ErrInvalidInput
	}

	// Все символы, кроме разрешенных
	forbiddenChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZабвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ0123456789"

	for _, char := range forbiddenChars {
		if strings.ContainsRune(text, char) {
			return false, nil
		}
	}
	return true, nil
}

func Checker(text string) (string, error) {
	isMorse, err := isMorse(text)
	if err != nil {
		logger.Println(err)
		return "", err
	}

	if isMorse == true {
		return morse.ToText(text), nil
	} else {
		return morse.ToMorse(text), nil
	}
}
