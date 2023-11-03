package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Инициализируем пустой словарь
	temperatureGroups := make(map[int][]float64)

	// Обходим все значения температур
	for _, temp := range temperatures {
		// Округляем значение температуры до ближайшего меньшего десятка
		groupKey := int(math.Floor(temp/10.0)) * 10

		// Добавляем значение температуры в соответствующую группу
		temperatureGroups[groupKey] = append(temperatureGroups[groupKey], temp)
	}

	// Сортируем ключи групп по возрастанию
	var keys []int
	for key := range temperatureGroups {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Выводим группы температур
	for _, key := range keys {
		fmt.Printf("%d: %v\n", key, temperatureGroups[key])
	}
}
