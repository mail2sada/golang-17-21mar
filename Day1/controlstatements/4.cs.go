package main

import "fmt"

func main() {
	fmt.Println("Demo: Logical opearators.. ")

	a, b, c := 10, 20, 30

	if a > b && a > c {
		fmt.Println("a is biggest")
	} else if b > c {
		fmt.Println("b is biggest")
	} else {
		fmt.Println("c is biggest")
	}

	if (a>b && a >c) && (c < a && c <b) {

	}
}
