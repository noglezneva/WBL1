// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

//первый способ(просто арифметика)
func swapNumbers1(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

//второй способ(XOR - искл. ИЛИ)
func swapNumbers2(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

// третий способ
func swapNumbers3(a, b int) (int, int) {
	return b, a
}

// четвертый способ
func swapNumbers4(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20

	// Пример использования первого способа
	x, y = swapNumbers1(x, y)
	fmt.Println(x, y)

	// Пример использования второго способа
	x, y = swapNumbers2(x, y)
	fmt.Println(x, y)

	// Пример использования третьего способа
	x, y = swapNumbers3(x, y)
	fmt.Println(x, y)

	// Пример использования четвертого способа
	swapNumbers4(&x, &y)
	fmt.Println(x, y)
}
