package main

import (
	"fmt"
	"mymodule/math"
	"mymodule/math/geometry"
)

func main() {
	fmt.Println("Demo: Packages and Modules...")
	c := math.Add(10, 20)

	fmt.Println("Result is", c)

	d := geometry.GetArea(2.0)

	fmt.Println(d)

	area := geometry.GetAreaR(10, 10)

	fmt.Println(area)
}
