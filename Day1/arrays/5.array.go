package main

import "fmt"

func main() {
	fmt.Println("Demo: Arrays")

	var array = [...]int{1, 2, 3, 4, 5}

	fmt.Println(array)

	var array1 = [...]int{0: 100, 1000: 500}
	fmt.Println(array1)

	fmt.Println(len(array), len(array1))
}
