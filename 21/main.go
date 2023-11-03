// Реализовать паттерн «адаптер» на любом примере
package main

import "fmt"

// Предположим, у нас есть интерфейс Printer со следующими методами:
type Printer interface {
	Print(message string)
}

// А у нас также есть структура OldPrinter с методом PrintLegacy,
// который не соответствует интерфейсу Printer:

type OldPrinter struct{}

func (p OldPrinter) PrintLegacy(text string) {
	fmt.Println("Printing (Legacy):", text)
}

// Мы хотели бы использовать старый принтер OldPrinter вместо нового принтера Printer.
// Для этого мы можем создать адаптер, который реализует интерфейс Printer
// и делегирует вызовы метода Print новому методу PrintLegacy:

type PrinterAdapter struct {
	OldPrinter
}

func (pa PrinterAdapter) Print(message string) {
	pa.PrintLegacy(message)
}

// Теперь мы можем использовать адаптер PrinterAdapter для вызова метода Print у старого принтера:
func main() {
	printer := PrinterAdapter{}
	printer.Print("Hello, World!")
}

// При выполнении кода будет выведено "Printing (Legacy): Hello, World!",
// и мы успешно адаптировали старый принтер к интерфейсу Printer
// с помощью паттерна "адаптер".
