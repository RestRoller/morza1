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

	// Если строка состоит ТОЛЬКО из точек, тире и пробелов - это код Морзе
	if isMorseCode(trimmed) {
		// Конвертируем Морзе в текст (английский)
		result := morse.ToText(trimmed)
		return strings.TrimSpace(result), nil
	} else {
		// Иначе конвертируем текст (английский) в Морзе
		result := morse.ToMorse(trimmed)
		return strings.TrimSpace(result), nil
	}
}

func isMorseCode(input string) bool {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return false
	}

	// Проверяем каждый символ в строке
	for _, char := range trimmed {
		// Если символ НЕ точка, НЕ тире, НЕ пробел и НЕ перевод строки - это не код Морзе
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			return false
		}
	}

	return true
}
