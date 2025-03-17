package main

import "fmt"

func main() {
	fmt.Println("Demo slice....")

	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(slc); i++ {
		fmt.Println(slc[i])
	}

	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	slc = append(slc, 100)

	fmt.Println(len(slc))
	fmt.Println(cap(slc))

	fmt.Println(slc[10])

	for idx, val := range slc {
		fmt.Println(idx, val)
	}
}
