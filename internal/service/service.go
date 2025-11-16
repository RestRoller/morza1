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

	// Простая проверка: если есть русские буквы - это текст, иначе морзе
	for _, char := range trimmed {
		if (char >= 'А' && char <= 'Я') || (char >= 'а' && char <= 'я') {
			// Есть русские буквы - конвертируем в морзе
			result := morse.ToMorse(trimmed)
			return strings.TrimSpace(result), nil
		}
	}

	// Нет русских букв - считаем что это морзе и конвертируем в текст
	result := morse.ToText(trimmed)
	return strings.TrimSpace(result), nil
}
