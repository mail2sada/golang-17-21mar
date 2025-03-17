package main

import "fmt"

func main() {
	fmt.Println("Demo: cs if - else")

	var a, b = 100, 500

	if a > b {
		fmt.Println("a is big")
	} else {
		fmt.Println("b is big")
	}
}
