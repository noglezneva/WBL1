// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package main

import (
	"fmt"
	"sync"
)

func main() {
	// исходный массив
	numbers := []int{2, 4, 6, 8, 10}
	// создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Используем буферизированный канал для сбора результатов
	results := make(chan int, len(numbers))

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			square := x * x
			results <- square
		}(num)
	}
	// Закрываем канал результатов после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()
	// Читаем квадраты чисел из канала результатов и выводим в stdout
	for square := range results {
		fmt.Println(square)
	}

}

// // другой способ решения
// func main() {
// 	// Исходный массив
// 	numbers := []int{2, 4, 6, 8, 10}
// 	// Создаем WaitGroup для ожидания завершения всех горутин
// 	var wg sync.WaitGroup

// 	// Создаем мьютекс для синхронизации доступа к результатам
// 	var mutex sync.Mutex

// 	// Создаем слайс для хранения результатов
// 	results := make([]int, len(numbers))

// 	// Запускаем горутину для каждого числа
// 	for i, num := range numbers {
// 		wg.Add(1)
// 		go func(index, x int) {
// 			defer wg.Done()
// 			square := x * x
// 			// Блокируем доступ к результатам и сохраняем квадрат по соответствующему индексу
// 			mutex.Lock()
// 			results[index] = square
// 			mutex.Unlock()
// 		}(i, num)

// 	}
// 	// Ожидаем завершения всех горутин
// 	wg.Wait()
// 	// Выводим результаты в stdout
// 	for _, square := range results {
// 		fmt.Println(square)
// 	}

// }
