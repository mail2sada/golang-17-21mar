package main

import "fmt"

func Divide(a, b int) int {
	if b == 0 {
		panic("Unable to handle divide by zero:Custom panic...")
	}
	return a / b
}

func main() {
	fmt.Println("Demo panic")

	fmt.Println(Divide(10, 2))

	fmt.Println(Divide(10, 0))

}
