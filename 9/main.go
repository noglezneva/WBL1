// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.
package main

import (
	"fmt"
)

func main() {
	// Создаем каналы
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Горутина для записи чисел в первый канал
	go func() {
		numbers := []int{1, 2, 3, 4, 5} // Входные числа
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan) // Закрываем канал после записи всех чисел
	}()

	// Горутина для умножения чисел из первого канала и записи во второй канал
	go func() {
		for num := range inputChan {
			outputChan <- num * 2
		}
		close(outputChan) // Закрываем второй канал после обработки всех чисел
	}()

	// Выводим результаты из второго канала в stdout
	for result := range outputChan {
		fmt.Println(result)
	}
}
