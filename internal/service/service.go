package service

import (
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"strings"
)

// AutoDetectAndConvert автоматически определяет тип содержимого и конвертирует его
func AutoDetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", fmt.Errorf("empty input")
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

	// Разбиваем на строки
	lines := strings.Split(trimmed, "\n")
	morseLines := 0
	totalLines := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		totalLines++

		// Разбиваем строку на слова
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		isMorseLine := true
		for _, word := range words {
			// Каждое слово должно состоять только из точек и тире
			for _, char := range word {
				if char != '.' && char != '-' {
					isMorseLine = false
					break
				}
			}
			if !isMorseLine {
				break
			}
		}

		if isMorseLine {
			morseLines++
		}
	}

	// Если больше половины строк являются кодом Морзе, считаем что это код Морзе
	if totalLines > 0 && morseLines*2 >= totalLines {
		return true
	}

	// Дополнительная проверка для коротких текстов
	if totalLines == 1 {
		line := strings.TrimSpace(lines[0])
		words := strings.Fields(line)
		morseWords := 0

		for _, word := range words {
			isMorseWord := true
			for _, char := range word {
				if char != '.' && char != '-' {
					isMorseWord = false
					break
				}
			}
			if isMorseWord {
				morseWords++
			}
		}

		// Если все слова состоят из точек и тире - это код Морзе
		if len(words) > 0 && morseWords == len(words) {
			return true
		}
	}

	return false
}
