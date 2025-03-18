package main

import "fmt"

func main() {
	fmt.Println("Demo Slice of slice...")

	var sliceOfSlice = [][]int{{1, 2, 3, 4, 5, 6}, {1, 2, 3}, {1, 2, 3, 4, 5, 6, 7}}

	fmt.Println("Contents of slice:", sliceOfSlice)

	for _, val := range sliceOfSlice {
		for _, v := range val {
			fmt.Println(v)
		}
	}
}
