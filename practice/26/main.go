package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке
func hasUniqueCharacters(s string) bool {
	// Создаем map для отслеживания символов
	charMap := make(map[rune]bool)

	// Приводим строку к нижнему регистру, чтобы сделать способ регистронезависимым
	s = strings.ToLower(s)

	// Проходим по каждому символу в строке
	for _, char := range s {
		// Если символ уже есть в map, возвращаем false - строка содержит неуникальные символы
		if _, exists := charMap[char]; exists {
			return false
		}

		// Добавляем символ в map
		charMap[char] = true
	}

	// Если ни один символ не повторился - в строке одни уникальные символы, возвращаем true
	return true
}

func main() {
	// Тестовые строки
	testCases := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	// Проверка каждой строки
	for _, testCase := range testCases {
		result := hasUniqueCharacters(testCase)
		fmt.Printf("%s: %v\n", testCase, result)
	}
}
