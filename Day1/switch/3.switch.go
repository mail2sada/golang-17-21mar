package main

import "fmt"

func main() {
	fmt.Println("Demo: switch continued")

	a := 100

	switch {
	case a >= 0 && a <= 9:
		fmt.Println("Single digit")
	case a >= 10 && a <= 99:
		fmt.Println("Double digit")
	case a >= 100 && a <= 999:
		fmt.Println("Three digits")
	default:
		fmt.Println("Un handled...")
	}

	switch {
	case a >= 1:
		fmt.Println("1")
	case a >= 10:
		fmt.Println("2")
	case a >= 100:
		fmt.Println("3")
	}
}
