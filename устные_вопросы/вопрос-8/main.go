// 8. В чем разница make и new?

// В Go различают два встроенных функции для создания значений: make и new. Вот их различия:

// make используется для создания и инициализации сложных типов данных, таких как slice, map и channel.
// Он возвращает инициализированный экземпляр указанного типа. Пример использования make:
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// slice := make([]int, 0, 5)
	// myMap := make(map[string]int)
	// channel := make(chan int)

	// new используется для выделения памяти под указанный тип данных
	// и возвращает указатель на нулевое значение этого типа.
	// Пример использования new:

	// var ptr *int
	// ptr = new(int)

	// функция new возвращает указатель, а не значение.
	// Новое значение будет нулевым (для примитивных типов)
	// или нулевыми значениями полей (для структур).
	// Используя функцию new, вы получаете указатель на выделенную память,
	// в то время как make инициализирует непосредственно значение.

	p := new(Person)
	fmt.Println(p) // Выводит указатель на новый экземпляр Person: &{"" 0}

	p.Name = "Alice"
	p.Age = 25
	fmt.Println(p) // Выводит указатель на тот же экземпляр Person: &{"Alice" 25}

}

// Таким образом, make и new используются в разных ситуациях и
// возвращают разные результаты.
