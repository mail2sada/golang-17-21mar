package main

import "fmt"

func main() {
	fmt.Println("Demo: Arrays...")

	var array [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)
	fmt.Println(array[0])

	array[4] = 500

	fmt.Println(array[4])
	fmt.Println(array)

}
