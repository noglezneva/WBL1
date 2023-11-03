package main

import (
	"fmt"
)

// Используя операции побитового сдвига и побитового ИЛИ
func main() {
	var num int64
	var i uint

	// Запрос ввода числа и позиции бита от пользователя
	fmt.Print("Введите число: ")
	_, _ = fmt.Scan(&num)
	fmt.Print("Введите позицию бита: ")
	_, _ = fmt.Scan(&i)

	// Установка i-го бита в 1
	num |= 1 << i
	fmt.Println("Число после установки бита в 1:", num)

	// Установка i-го бита в 0
	num &= ^(1 << i)
	fmt.Println("Число после установки бита в 0:", num)
}