package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {

	point := NewPoint(1.2, 2.5)
	point1 := NewPoint(5, 8.6)

	fmt.Println(distance(point, point1))

}

func distance(point, point1 *Point) float64 {
	return math.Sqrt(math.Pow((point.x-point1.x), 2) + math.Pow((point.y-point1.y), 2))
}