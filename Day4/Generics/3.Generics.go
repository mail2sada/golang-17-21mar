package main

import "fmt"

func Swap[T any](a, b T) (T, T) {
	// c := a
	// a = b
	// b = c
	// return a, b
	return b, a
}

type tmp struct {
	x, y int
}

func main() {
	fmt.Println("Generic Swap")

	a, b := Swap("Hi", "Hello")

	fmt.Println(a, b)

	c, d := Swap(1, 2)

	fmt.Println(c, d)

	e, f := Swap(tmp{x: 10, y: 5}, tmp{x: 100, y: 50})

	fmt.Println(e, f)

}
