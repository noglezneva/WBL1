// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}

package main

import (
	"fmt"
	"reflect"
)

// первый способ
func getType(value interface{}) string {
	// Получаем тип переменной в виде reflect.Type
	t := reflect.TypeOf(value)

	// Возвращаем строковое представление типа
	return t.String()
}

// второй способ
func getTypeSwitch(value interface{}) string {
	// Используем switch для определения типа переменной
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "channel"
	default:
		return "unknown"
	}
}

func main() {
	// Примеры переменных разных типов
	var num int = 42
	var name string = "Hello, world!"
	var flag bool = true
	var ch chan int = make(chan int)

	// Определение типа переменных в рантайме
	numType := getType(num)
	nameType := getType(name)
	flagType := getType(flag)
	chType := getType(ch)

	// Вывод типов переменных
	fmt.Println("Type of num:", numType)
	fmt.Println("Type of name:", nameType)
	fmt.Println("Type of flag:", flagType)
	fmt.Println("Type of ch:", chType)

	// // Определение типа переменных в рантайме с использованием switch
	// numType := getTypeSwitch(num)
	// nameType := getTypeSwitch(name)
	// flagType := getTypeSwitch(flag)
	// chType := getTypeSwitch(ch)

	// // Вывод типов переменных
	// fmt.Println("Type of num:", numType)
	// fmt.Println("Type of name:", nameType)
	// fmt.Println("Type of flag:", flagType)
	// fmt.Println("Type of ch:", chType)
}
