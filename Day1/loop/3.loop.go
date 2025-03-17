package main

import "fmt"

func main() {
	fmt.Println("Demo: loops")

	i := 10

	for i > 0 {
		fmt.Println("i=", i)
		i--
	}
}
