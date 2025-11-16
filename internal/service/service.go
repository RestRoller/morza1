package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertData(input string) (string, error) {
	cleanedInput := strings.TrimSpace(input)
	if cleanedInput == "" {
		return "", nil
	}

	if isMorseEncoding(cleanedInput) {
		result := morse.ToText(cleanedInput)
		return strings.TrimSpace(result), nil
	}

	result := morse.ToMorse(cleanedInput)
	return strings.TrimSpace(result), nil
}

func isMorseEncoding(data string) bool {
	trimmed := strings.TrimSpace(data)
	if trimmed == "" {
		return false
	}

	for _, char := range trimmed {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			return false
		}
	}
	return true
}
