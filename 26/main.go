// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
package main

import (
	"fmt"
	"strings"
)

// Функция IsUnique проверяет, являются ли все символы в строке уникальными
func IsUnique(str string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	str = strings.ToLower(str)

	// Создаем мапу для отслеживания уникальных символов
	seen := make(map[rune]bool)

	// Проходимся по всем символам строки
	for _, char := range str {
		// Если символ уже встречался ранее, значит он не уникален
		if seen[char] {
			return false
		}

		// Добавляем символ в мапу, чтобы отметить его как уникальный
		seen[char] = true
	}

	// Все символы в строке уникальны
	return true
}

func main() {
	str := "abcd"
	fmt.Printf("Строка \"%s\" содержит уникальные символы: %v\n", str, IsUnique(str))

	str = "abCdefAaf"
	fmt.Printf("Строка \"%s\" содержит уникальные символы: %v\n", str, IsUnique(str))

	str = "aabcd"
	fmt.Printf("Строка \"%s\" содержит уникальные символы: %v\n", str, IsUnique(str))
}
