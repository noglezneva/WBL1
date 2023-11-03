// Удалить i-ый элемент из слайса
package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Исходный слайс 1: ", slice1)
	fmt.Println("Исходный слайс 2: ", slice2)
	i := 2
	i2 := 5

	// Удаляем i-ый элемент из слайса
	slice1 = append(slice1[:i], slice1[i+1:]...)
	slice2 = deleteElement2(slice2, i2)

	fmt.Println("1 способ: Слайс после удаление по индексу ", i, ":", slice1)
	fmt.Println("2 способ: Слайс после удаление по индексу ", i2, ":", slice2)

}

// способ 2 (используя копирование срезов)
func deleteElement2(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
