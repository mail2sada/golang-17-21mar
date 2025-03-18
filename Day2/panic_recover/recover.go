package main

import "fmt"

func divide(a, b int) int {
	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("handled panic")
		}
	}()

	return a / b
}

func main() {
	fmt.Println("Demo panic handling")

	defer func() {
		a := recover()
		if a != nil {
			fmt.Println("Received panic, can't handle...")
		}
	}()

	fmt.Println(divide(10, 1))

	fmt.Println("Exiting main...")
}
