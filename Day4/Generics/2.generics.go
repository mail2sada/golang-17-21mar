package main

import "fmt"

func Display[T any](elements ...T) {
	for idx, v := range elements {
		fmt.Println(idx, v)
	}
}

func Print(elements ...interface{}) {
	for idx, v := range elements {
		fmt.Println(idx, v)
	}
}

type temp struct {
	a, b int
}

func main() {
	fmt.Println("Constraints...")

	Display("Hello", "Hi", "Test")

	Display(1, 2, 3, 4, 5)

	Display(temp{a: 10, b: 5}, temp{a: 4})
	Display(0)
	Print()
	Print(1, "Hello", temp{a: 10, b: 5})

	Print(1, 2)
}
