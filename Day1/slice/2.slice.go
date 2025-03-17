package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice")

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
