//5
// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

package main

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i                 //Отправляем значения в канал
		time.Sleep(time.Second) // Задержка 1 секунда
	}
	close(ch) //Закрываем канал после отправки всех значений
}

func receiveData(ch <-chan int) {
	for value := range ch {
		fmt.Println(value) //Читаем значения из канала
	}
}

func main() {
	N := 5               //кол-во секунд работы программы
	ch := make(chan int) // Создаем канал
	go sendData(ch)      // Запускаем горутину для отправки данных
	go receiveData(ch)   // Запускаем горутину для чтения данных

	time.Sleep(time.Duration(N) * time.Second) // Задержка на N секунд

	fmt.Println("Программа завершена!")

}
