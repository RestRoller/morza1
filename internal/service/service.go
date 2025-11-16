package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", nil
	}

	if isMorseCode(trimmed) {
		result := morse.ToText(trimmed)
		return strings.TrimSpace(result), nil
	}

	result := morse.ToMorse(trimmed)
	return strings.TrimSpace(result), nil
}

func isMorseCode(input string) bool {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return false
	}

	// Проверяем каждый символ
	for _, char := range trimmed {
		// Если нашли любой символ кроме точек, тире и пробелов - это не морзе
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			return false
		}
	}
	return true
}
