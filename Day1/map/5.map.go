package main

import "fmt"

func main() {
	fmt.Println("Demo: deleting element in map")

	mp := map[float32]int{10.1: 100, 20.2: 200, 30.5: 400}

	fmt.Println(mp)

	for key, val := range mp {
		fmt.Println(key, val)
	}

	delete(mp, 10.1)
	fmt.Println(mp)

}
