//9. Сколько существует способов задать переменную типа slice или map?

package main

import "fmt"

func main() {
	// 1. Инициализация с использованием литералов:
	// Инициализация slice
	mySlice := []int{1, 2, 3, 4, 5}

	// Инициализация map
	myMap := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	fmt.Println(mySlice, myMap)

	//2. Использование функции make:
	// Инициализация slice
	mySlice2 := make([]int, 5)

	// Инициализация map
	myMap2 := make(map[string]int)
	fmt.Println(mySlice2, myMap2)

	//3 Использование пустого литерала:

	// Инициализация пустого slice
	var mySlice3 []int

	// Инициализация пустого map
	var myMap3 map[string]int
	fmt.Println(mySlice3, myMap3)

	//4 Присваивание значения существующей переменной типа slice или map:
	var mySlice4 []int
	mySlice4 = []int{1, 2, 3}

	// Присваивание значения map
	var myMap4 map[string]int
	myMap4 = map[string]int{
		"apple":  1,
		"banana": 2,
	}
	fmt.Println(mySlice4, myMap4)
}
