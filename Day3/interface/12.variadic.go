package main

import "fmt"

func Add(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func main() {
	sum := Add()

	fmt.Println(sum)

	sum = Add(1)
	fmt.Println(sum)

	sum = Add(1, 2)

	fmt.Println(sum)

	sum = Add(1, 2, 3)
	fmt.Println(sum)

	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum = Add(slc...)
	fmt.Println(sum)
}
