package service

import (
	"strings"
)

// Руская таблица Морзе
var russianMorse = map[rune]string{
	'А': ".-", 'Б': "-...", 'В': ".--", 'Г': "--.", 'Д': "-..",
	'Е': ".", 'Ё': ".", 'Ж': "...-", 'З': "--..", 'И': "..",
	'Й': ".---", 'К': "-.-", 'Л': ".-..", 'М': "--", 'Н': "-.",
	'О': "---", 'П': ".--.", 'Р': ".-.", 'С': "...", 'Т': "-",
	'У': "..-", 'Ф': "..-.", 'Х': "....", 'Ц': "-.-.", 'Ч': "---.",
	'Ш': "----", 'Щ': "--.-", 'Ъ': "--.--", 'Ы': "-.--", 'Ь': "-..-",
	'Э': "..-..", 'Ю': "..--", 'Я': ".-.-",
}

var morseRussian = map[string]rune{
	".-": 'А', "-...": 'Б', ".--": 'В', "--.": 'Г', "-..": 'Д',
	".": 'Е', "...-": 'Ж', "--..": 'З', "..": 'И', ".---": 'Й',
	"-.-": 'К', ".-..": 'Л', "--": 'М', "-.": 'Н', "---": 'О',
	".--.": 'П', ".-.": 'Р', "...": 'С', "-": 'Т', "..-": 'У',
	"..-.": 'Ф', "....": 'Х', "-.-.": 'Ц', "---.": 'Ч', "----": 'Ш',
	"--.-": 'Щ', "--.--": 'Ъ', "-.--": 'Ы', "-..-": 'Ь', "..-..": 'Э',
	"..--": 'Ю', ".-.-": 'Я',
}

func AutoDetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", nil
	}

	if isMorseCode(trimmed) {
		return morseToRussian(trimmed), nil
	} else {
		return russianToMorse(trimmed), nil
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

func russianToMorse(text string) string {
	text = strings.ToUpper(text)
	var result []string

	for _, char := range text {
		if char == ' ' {
			result = append(result, " ")
			continue
		}
		if morse, exists := russianMorse[char]; exists {
			result = append(result, morse)
		} else {
			result = append(result, string(char))
		}
	}

	return strings.Join(result, " ")
}

func morseToRussian(morse string) string {
	words := strings.Split(morse, "   ")
	var result []string

	for _, word := range words {
		codes := strings.Split(word, " ")
		var wordChars []string

		for _, code := range codes {
			if char, exists := morseRussian[code]; exists {
				wordChars = append(wordChars, string(char))
			} else if code != "" {
				wordChars = append(wordChars, code)
			}
		}

		result = append(result, strings.Join(wordChars, ""))
	}

	return strings.Join(result, " ")
}
