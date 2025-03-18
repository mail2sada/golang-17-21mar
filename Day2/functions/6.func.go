package main

import "fmt"

func main() {
	fmt.Println("Named returns...")
	a, b := PropSquare(10)

	fmt.Println(a, b)
}

func PropSquare(side float64) (area, circ float64) {
	area = side * side
	circ = 4 * side
	return
}
