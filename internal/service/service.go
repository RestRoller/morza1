package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func detectMorseEncoding(inputData string) bool {
	cleanData := strings.TrimSpace(inputData)
	if cleanData == "" {
		return false
	}

	for _, char := range cleanData {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\t' && char != '\r' {
			return false
		}
	}
	return true
}

func ConvertContent(inputContent string) string {
	trimmedContent := strings.TrimSpace(inputContent)
	if trimmedContent == "" {
		return ""
	}

	if detectMorseEncoding(trimmedContent) {
		decodedText := morse.ToText(trimmedContent)
		return strings.TrimSpace(decodedText)
	} else {
		encodedMorse := morse.ToMorse(trimmedContent)
		return strings.TrimSpace(encodedMorse)
	}
}
