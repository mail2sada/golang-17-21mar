package main

import "fmt"

func main() {
	fmt.Println("defer in detail")

	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a, b := 10, 20

	if a < b {

		defer fmt.Println("a is small")
	}

	for i := 0; i < len(slc); i++ {
		defer fmt.Println(slc[i])
	}

	fmt.Println("Exiting main..")
}
