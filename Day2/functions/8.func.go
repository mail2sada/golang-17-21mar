package main

import "fmt"

func ReadMySlice(slice []int, res func(int)) {
	for _, v := range slice {
		res(v)
	}
}

func main() {
	fmt.Println("Demo: Annonymous functions")

	a := func(a int) {
		fmt.Println(a * 2)
	}

	ReadMySlice([]int{10, 20, 30, 40, 50}, a)
}
