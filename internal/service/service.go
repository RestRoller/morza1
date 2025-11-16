package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

// AutoDetectAndConvert автоматически определяет тип содержимого и конвертирует его
func AutoDetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", nil
	}

	if isMorseCode(trimmed) {
		// Если это код Морзе, конвертируем в текст
		result := morse.ToText(trimmed)
		return strings.TrimSpace(result), nil
	}
	// Если это текст, конвертируем в код Морзе
	result := morse.ToMorse(trimmed)
	return strings.TrimSpace(result), nil
}

// isMorseCode проверяет, является ли строка кодом Морзе
func isMorseCode(input string) bool {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return false
	}

	// Считаем количество символов Морзе (точки и тире)
	morseCount := 0
	totalCount := 0

	for _, char := range trimmed {
		if char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			totalCount++
			if char == '.' || char == '-' {
				morseCount++
			}
		}
	}

	// Если более 90% символов - точки или тире, считаем что это код Морзе
	if totalCount > 0 && float64(morseCount)/float64(totalCount) >= 0.9 {
		return true
	}

	return false
}
