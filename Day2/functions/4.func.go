package main

import "fmt"

func PropertiesOfCircle(radius float64) (float64, float64) {
	pi := 3.14159
	area := pi * radius * radius
	circ := 2 * pi * radius
	return area, circ
}

func main() {
	fmt.Println("Demo: returning multiple values")

	a, c := PropertiesOfCircle(10.1)

	fmt.Println(a, c)

	_, c = PropertiesOfCircle(10.5)

	fmt.Println(c)

	a, _ = PropertiesOfCircle(5.5)

	fmt.Println(a)
}
