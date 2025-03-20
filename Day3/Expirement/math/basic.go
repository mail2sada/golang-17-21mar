package math

import "fmt"

func init() {
	fmt.Println("Init from math")
}

func init() {
	fmt.Println("Init2 from math ")
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
