// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
package main

import (
	"fmt"
	"math"
)

// Определение структуры Point
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Метод для расчета расстояния между двумя точками
func (p Point) DistanceTo(q Point) float64 {
	dx := q.x - p.x
	dy := q.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создание двух точек
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	// Расчет расстояния между точками
	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
