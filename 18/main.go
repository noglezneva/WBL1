// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговое
// значение счетчика.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func main() {
	// Создаем экземпляр структуры Counter
	counter := Counter{}

	// Задаем количество горутин для инкрементирования счетчика
	numWorkers := 10

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запускаем горутины для инкрементирования счетчика
	for i := 0; i < numWorkers; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.value)
}
