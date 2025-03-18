package main

import "fmt"

func main() {
	fmt.Println("Demo: Strings")

	var str string = "Welcome to go training day2"

	fmt.Println(str)

	// Explore slice of byte analogy

	for i, v := range str {
		fmt.Println(i, v)
	}

	str2 := "Hello, 世界"

	fmt.Println(len(str2))

	for i, v := range str2 {
		fmt.Println(i, v)
	}
}
