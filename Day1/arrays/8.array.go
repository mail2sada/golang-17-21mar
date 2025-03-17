package main

import "fmt"

func main() {
	fmt.Println("Range operator...")

	array := [5]int{1, 2, 3, 4, 5}

	for idx, val := range array {
		fmt.Println(idx, val)

		val *= idx
	}

	for idx, val := range array {
		fmt.Println(idx, val)
	}
}
