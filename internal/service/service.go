package service

import (
	"errors"
	"strings"

	"github.com/maximakhmatovich/go_http-final/pkg/morse"
)

var ErrInvalidInput = errors.New("Неверные входные данные")

func isMorse(text string) (bool, error) {
	if text == "" {
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
		return "", err
	}

	if isMorse == true {
		return morse.ToText(text), nil
	} else {
		return morse.ToMorse(text), nil
	}
}
