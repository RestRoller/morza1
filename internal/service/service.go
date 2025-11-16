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

	if totalCount == 0 {
		return false
	}

	return float64(morseCount)/float64(totalCount) >= 0.9
}
