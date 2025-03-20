package main

import "fmt"

type Shape interface {
	Area() float64
	Perimiter() float64
	Sides() int
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}
func (c Circle) Perimiter() float64 {
	return 2 * 3.14159 * c.Radius
}
func (c Circle) Sides() int {
	return -1
}

type Rectangle struct {
	Length, Breadth int
}

func (c Rectangle) Area() float64 {
	return float64(c.Length) * float64(c.Breadth)
}
func (c Rectangle) Perimiter() float64 {
	return float64(2 * (c.Length + c.Breadth))
}
func (c Rectangle) Sides() int {
	return 4
}

func main() {

	fmt.Println("Interfaces...")

	c := Circle{Radius: 1.0}

	r := Rectangle{Length: 2.0, Breadth: 1.0}

	fmt.Println(c.Area())
	fmt.Println(c.Perimiter())
	fmt.Println(c.Sides())

	fmt.Println(r.Area())
	fmt.Println(r.Perimiter())
	fmt.Println(r.Sides())

	var s Shape

	s = c

	fmt.Println(s.Area())

	s = r

	fmt.Println(s.Sides())

	fmt.Println(s)
}
