package main

import "fmt"

func main() {
	fmt.Println("Demo: range operator")

	array := [5]int{10, 20, 30, 40, 50}

	for idx, _ := range array {
		fmt.Println(idx)
	}
}
