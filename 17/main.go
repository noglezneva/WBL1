// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Отсортированный массив, в котором будем искать
	nums := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Значение, которое хотим найти
	target := 9

	// Используем встроенный метод sort.Search() для бинарного поиска
	index := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	// Проверяем, найдено ли значение
	if index < len(nums) && nums[index] == target {
		fmt.Printf("Значение %d найдено на позиции %d\n", target, index)
	} else {
		fmt.Printf("Значение %d не найдено\n", target)
	}
}
