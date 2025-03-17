package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice[:3])

	fmt.Println(slice[3:])

	idx := 4

	fmt.Println(slice[0 : idx+1])

	fmt.Println(slice[idx:])

	slice = append(slice[0:idx+1], slice[idx:]...)

	fmt.Println(slice)

	slice[idx] = 1000
	fmt.Println(slice)

}
