package main

import "fmt"

func main() {
	fmt.Println("Demo:Slice deleting a element")

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	idx := 4

	fmt.Println(slice[:idx])
	fmt.Println(slice[idx+1:])

	slice = append(slice[:idx], slice[idx+1:]...)
	fmt.Println(slice)
}
