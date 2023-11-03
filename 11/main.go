// Реализовать пересечение двух неупорядоченных множеств.

// Для реализации пересечения двух неупорядоченных множеств на языке Go
// можно воспользоваться стандартным пакетом sort для сортировки элементов
// в множествах и алгоритмом "двух указателей" (two-pointer algorithm).

package main

import (
	"fmt"
	"sort"
)

func intersection(set1, set2 []int) []int {
	// Сортируем множества
	sort.Ints(set1)
	sort.Ints(set2)

	var result []int

	// Инициализируем два указателя для обхода множеств
	i, j := 0, 0

	for i < len(set1) && j < len(set2) {
		if set1[i] == set2[j] {
			// Если элементы равны, добавляем его в результат и двигаем оба указателя
			result = append(result, set1[i])
			i++
			j++
		} else if set1[i] < set2[j] {
			// Если элемент в set1 меньше, двигаем указатель set1
			i++
		} else {
			// Если элемент в set2 меньше, двигаем указатель set2
			j++
		}
	}

	return result
}

func main() {
	set1 := []int{5, 2, 9, 12, 10}
	set2 := []int{2, 7, 5, 15}

	intersected := intersection(set1, set2)

	fmt.Println(intersected) // Выводит [2, 5]
}
