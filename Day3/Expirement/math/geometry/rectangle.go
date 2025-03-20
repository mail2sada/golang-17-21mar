package geometry

import "fmt"

func init() {
	fmt.Println("Hello from rectangle")
}

type Rectangle struct {
	Lenght, Breadth float64
}

func (r Rectangle) GetArea() float64 {
	return r.Lenght * r.Breadth
}

func (r Rectangle) GetPerimeter() float64 {
	return 2 * (r.Lenght + r.Breadth)
}
