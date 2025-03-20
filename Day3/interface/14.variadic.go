package main

import "fmt"

func add(a, b int, ints ...int) int {
	sum := a + b
	for _, v := range ints {
		sum += v
	}
	return sum
}

func main() {
	sum := 0

	sum = add(1, 2)

	fmt.Println(sum)

	sum = add(1, 2, 3)
	fmt.Println(sum)

	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum = add(10, 20, slc...)
	fmt.Println(sum)
}
