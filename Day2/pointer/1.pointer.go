package main

import "fmt"

func main() {
	fmt.Println("Demo: Pointers...")

	var i int = 10

	var ptr *int = &i

	var ptrptr **int = &ptr

	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 100

	fmt.Println(i)

	var ptrArray *[5]int

	var ptrSlice *[]int
}
