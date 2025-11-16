package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(s string) bool {
	clean := strings.TrimSpace(s)
	if clean == "" {
		return false
	}

	morseCount := 0
	totalCount := 0

	for _, char := range clean {
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

func ToggleMorse(s string) string {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return ""
	}

	if isMorse(trimmed) {
		result := morse.ToText(trimmed)
		return strings.TrimSpace(result)
	}
	result := morse.ToMorse(trimmed)
	return strings.TrimSpace(result)
}
