package main

import "fmt"

func main() {
	fmt.Println("Multiple returns..")

	a, c, d := PropCircle(10.5)

	fmt.Println(a, c, d)

	a, _, d = PropCircle(11.5)

	fmt.Println(a, d)
}

func PropCircle(radius float64) (float64, float64, float64) {

	pi := 3.14159

	area := pi * radius * radius

	cir := 2 * pi * radius

	dia := 2 * radius

	return area, cir, dia
}
