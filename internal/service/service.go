package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

// AutoDetectAndConvert автоматически определяет тип содержимого и конвертирует его
func AutoDetectAndConvert(input string) (string, error) {
	if isMorseCode(input) {
		// Если это код Морзе, конвертируем в текст
		return morse.ToText(input), nil
	}
	// Если это текст, конвертируем в код Морзе
	return morse.ToMorse(input), nil
}

// isMorseCode проверяет, является ли строка кодом Морзе
func isMorseCode(input string) bool {
	// Убираем лишние пробелы по краям
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return false
	}

	// Разбиваем на строки для анализа
	lines := strings.Split(trimmed, "\n")

	// Анализируем первые несколько строк для определения типа
	for i := 0; i < len(lines) && i < 3; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// Разбиваем строку на слова
		words := strings.Fields(line)
		for _, word := range words {
			// Код Морзе состоит только из точек, тире и пробелов
			for _, char := range word {
				if char != '.' && char != '-' {
					return false
				}
			}
		}
	}

	return true
}
