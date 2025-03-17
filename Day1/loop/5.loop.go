package main

import "fmt"

func main() {
	fmt.Println("Demo: loops")

	i := 10

	for {
		fmt.Println(i)
		if i >= 10000 {
			break
		}
		i++
	}
}
