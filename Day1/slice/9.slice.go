package main

import "fmt"

func main() {
	fmt.Println("Demo slice of arrays")

	sliceOfArray := [][4]int{0: {1, 2, 3, 4}, {1, 2, 3, 4}, {3, 4, 5, 6}, 5: {3, 1, 2}}

	fmt.Println(sliceOfArray)
}
