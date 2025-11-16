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

	for _, char := range clean {
		// Если нашли любой символ кроме точек, тире и пробелов - это не морзе
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			return false
		}
	}
	return true
}

func ToggleMorse(s string) string {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return ""
	}

	if isMorse(trimmed) {
		return strings.TrimSpace(morse.ToText(trimmed))
	}
	return strings.TrimSpace(morse.ToMorse(trimmed))
}
