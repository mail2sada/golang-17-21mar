package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) GetDiameter() float64 {
	return 2 * c.Radius
}

func (c Circle) GetArea() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) GetCircumfrance() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c Circle) PrintCircleProperties() {
	fmt.Println("Diameter:", c.GetDiameter())
	fmt.Println("Circumfrence:", c.GetCircumfrance())
	fmt.Println("Area:", c.GetArea())
}

func main() {
	fmt.Println("Demo: Methods")

	circle := Circle{Radius: 2.25}

	// d := circle.GetDiameter()

	// c := circle.GetCircumfrance()

	// a := circle.GetArea()

	// fmt.Println(d, c, a)

	circle.PrintCircleProperties()
}
