// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования)

package main

import "fmt"

// Определяем структуру Human
type Human struct {
	name    string
	age     int
	country string
}

// Определяем метод GetName для структуры Human
func (h *Human) GetName() string {
	return h.name
}

// Определяем метод GetAge для структуры Human
func (h *Human) GetAge() int {
	return h.age
}

// Определяем метод GetCountry для структуры Human
func (h *Human) GetCountry() string {
	return h.country
}

// Определяем структуру Action с встроенной структурой Human
type Action struct {
	Human
}

// Определяем новые методы, встроенные в структуру Action
func (a *Action) Walk() {
	fmt.Printf("%s is walking.\n", a.name) // Доступ к полю name из структуры Human
}

func (a *Action) Speak(message string) {
	fmt.Printf("%s is speaking: %s\n", a.name, message)
}

func main() {
	//Создаем экземпляр структуры Action
	person := Action{
		Human: Human{
			name:    "Nadia",
			age:     21,
			country: "Russia",
		},
	}
	//Обращаемся к методам из родительской структуры Human
	fmt.Println(person.GetName())    //Nadia
	fmt.Println(person.GetAge())     //21
	fmt.Println(person.GetCountry()) //Russia

	// Вызываем новый методы, встроенные в структуру Action
	person.Walk()          //Nadia is walking.
	person.Speak("Привет") //Nadia is speaking: Привет
}
