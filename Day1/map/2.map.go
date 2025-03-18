package main

import "fmt"

func main() {
	fmt.Println("Demo maps")

	mp := map[string][]int{"sdc": {1, 2, 3}, "wrg": {4, 5, 6}, "tas": {7, 8, 9}}

	mp["mco"] = []int{10, 11, 12}

	fmt.Println("Contents mco", mp["mco"])

}
