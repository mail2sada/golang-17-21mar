package main

import "fmt"

func main() {
	fmt.Println("Arrays...")

	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
		array[i] *= i // modifying array element
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
