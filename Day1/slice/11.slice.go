package main

import "fmt"

type temp struct {
}

func main() {
	fmt.Println("Demo slice with make")

	slice := make([]int, 10, 100)

	sliceArray := make([][]int, 0)

	sliceArray1 := make([][5]int, 0)

	fmt.Println(sliceArray)

	fmt.Println(sliceArray1)

	fmt.Println(len(slice))

	fmt.Println(cap(slice))

	fmt.Println(slice)

	for _, v := range slice {
		fmt.Println(v)
	}
}
