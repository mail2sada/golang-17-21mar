package main

import "fmt"

func main() {
	fmt.Println("Demo: slice...")

	var slice = []int{}

	fmt.Println(slice)

	fmt.Println(len(slice))

	fmt.Println(cap(slice))
}
