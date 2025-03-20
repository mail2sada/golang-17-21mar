package geometry

import "fmt"

func init() {
	fmt.Println("Init from geometry")
}

type Circle struct {
	Radius float64
}

const PI = 3.14159

func (c Circle) GetArea() float64 {
	return PI * c.Radius * c.Radius
}

func (c Circle) GetCircumfrance() float64 {
	return 2 * PI * c.Radius
}
