// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.
// var justString string
// func someFunc() {
// v := createHugeString(1 << 10)
// justString = v[:100]
// }
// func main() {
// someFunc()
// }
//Ответ:
// В данном фрагменте кода есть потенциальная проблема с утечкой памяти.
// При вызове функции createHugeString(1 << 10) создаётся огромная строка и присваивается переменной v.
// Затем первые 100 символов строки v копируются в переменную justString. Однако, оригинальная огромная строка v не будет освобождена,
// и это может привести к потреблению большого объема памяти, если это будет повторяться много раз.
//Чтобы исправить эту проблему, можно использовать bytes.Buffer, который автоматически увеличивает свой размер по мере необходимости.
//Корректный пример реализации:

package main

import (
	"bytes"
	"fmt"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	var buffer bytes.Buffer
	buffer.WriteString(v[:100])
	justString = buffer.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(size int) string {
	var buffer bytes.Buffer
	for i := 0; i < size; i++ {
		buffer.WriteString("x")
	}
	return buffer.String()
}

// В этом примере мы используем буфер bytes.Buffer, чтобы постепенно формировать строку justString.
// Мы вызываем функцию createHugeString, которая создает огромную строку заданного размера.
// Затем мы создаем новый буфер bytes.Buffer и записываем первые 100 символов строки v в буфер.
// Наконец, мы присваиваем строковое представление буфера justString.
//Таким образом, мы избегаем копирования полной огромной строки и предотвращаем утечку памяти.
