// Реализовать конкурентную запись данных в map.
// использую мьютекс для реализации доступа к map
package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]string
}

// метод для установки ключей и значений
func (sm *SafeMap) Set(key, value string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

// метод для извлечения значений
func (sm *SafeMap) Get(key string) string {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.m[key]
}

func main() {
	sm := SafeMap{
		m: make(map[string]string),
	}
	var wg sync.WaitGroup
	numWorkers := 5

	// Запуск нескольких горутин для конкурентной записи данных в map

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", workerID)
			value := fmt.Sprintf("value%d", workerID)
			sm.Set(key, value)
			fmt.Printf("Worker %d записал данные: key=%s, value=%s\n", workerID, key, value)
		}(i)
	}
	wg.Wait()

	//Вывод значений из map
	for i := 0; i < numWorkers; i++ {
		key := fmt.Sprintf("key%d", i)
		value := sm.Get(key)
		fmt.Printf("Значение из map: key=%s, value=%s\n", key, value)
	}
}

// Горутины вызывают метод Set для установки ключей и значений. Затем, значения извлекаются с помощью метода Get и выводятся на экран.
// Путем синхронизации доступа к map с помощью мьютекса,
// мы обеспечиваем конкурентную запись данных без возникновения гонок данных или других проблем,
// связанных с параллельным доступом к map.
