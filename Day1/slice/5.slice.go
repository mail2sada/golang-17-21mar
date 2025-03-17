package main

import "fmt"

func main() {
	fmt.Println("Operations with slice..")

	slc := []int{}

	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	slc = append(slc, 1)

	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	slc = append(slc, 2)

	fmt.Println(len(slc))
	fmt.Println(cap(slc))
}
