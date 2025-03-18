package main

import "fmt"

type temp struct {
	x, y int
}

func main() {
	fmt.Println("Slice of structures...")

	sliceOfstructs := []temp{temp{x: 1, y: 2}, temp{x: 10, y: 20}}

	fmt.Println(sliceOfstructs)

	for _, v := range sliceOfstructs {
		fmt.Println("X:", v.x)
		fmt.Println("Y:", v.y)
	}
}
