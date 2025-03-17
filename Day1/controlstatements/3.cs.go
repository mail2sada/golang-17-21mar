package main

import "fmt"

func main() {
	fmt.Println("Demo: CS if else if")

	a, b := 100, 100

	if a > b {
		fmt.Println("a is big")
	} else if a < b {
		fmt.Println("b is big")
	} else {
		fmt.Println("Looks like both are equal...")
	}
}
