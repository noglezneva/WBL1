// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
package main

import (
	"fmt"
)

func reverseString(inputString string) string {
	runes := []rune(inputString)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Print("Введите строку для переворота: ")
	var inputString string
	fmt.Scanln(&inputString)
	reversedString := reverseString(inputString)
	fmt.Println("Перевернутая строка:", reversedString)
}
