package main

import "fmt"

func main() {

	fmt.Println("Demo: Multidimensional array")

	mArray := [5][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}

	fmt.Println(mArray)

	for i := 0; i < len(mArray); i++ {
		for j := 0; j < len(mArray[i]); j++ {
			fmt.Println(mArray[i][j])
		}
	}

	for i, val := range mArray {

		for j, v := range val {
			fmt.Println(v, i, j)
		}

	}
}
