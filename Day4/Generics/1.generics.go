package main

import "fmt"

type Numbers interface {
	int | float64 | uint | uint64
}

func Add[T Numbers](a, b T) T {
	return a + b
}

func Sub[T Numbers](a, b T) T {
	return a - b
}

func main() {
	fmt.Println("Working on Generics...")

	c := Add(10, 20)

	fmt.Println("Result is", c)

	d := Add(10.5, 11.5)

	fmt.Println("Result is", d)

	e := Sub(10, 5)

	fmt.Println(e)

	e1 := Sub(10.5, 3.5)
	fmt.Println(e1)

}
