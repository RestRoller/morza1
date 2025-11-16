package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", nil
	}

	fmt.Printf("DEBUG: Input: '%s'\n", trimmed)
	fmt.Printf("DEBUG: IsMorse: %v\n", isMorseCode(trimmed))

	if isMorseCode(trimmed) {
		result := morse.ToText(trimmed)
		fmt.Printf("DEBUG: Morse->Text: '%s'\n", result)
		return strings.TrimSpace(result), nil
	} else {
		result := morse.ToMorse(trimmed)
		fmt.Printf("DEBUG: Text->Morse: '%s'\n", result)
		return strings.TrimSpace(result), nil
	}
}

func isMorseCode(input string) bool {
	trimmed := strings.TrimSpace(input)
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
