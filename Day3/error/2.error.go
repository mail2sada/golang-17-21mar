package main

import (
	"errors"
	"fmt"
)

func GetElement(slice []int, index int) (ans int, err error) {
	if index >= len(slice) {
		err = errors.New("Index out of bounds...")
		return
	}
	ans = slice[index]
	return
}

func main() {
	fmt.Println("Demo: Errors")

	ans, err := GetElement([]int{1, 2, 3, 4, 5, 6}, 1)

	if err != nil {
		fmt.Println("Received error...")
	} else {
		fmt.Println("value at index 1", ans)
	}
}
