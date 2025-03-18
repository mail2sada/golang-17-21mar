package main

import "fmt"

func main() {
	fmt.Println("Demo resetting map")

	mp := map[int]int{0: 1, 1: 2, 2: 3, 3: 4}

	fmt.Println(mp)

	for key, _ := range mp {
		delete(mp, key)
	}

	fmt.Println(mp)

	mp = map[int]int{0: 1, 1: 2, 2: 3, 3: 4}

	fmt.Println(mp)

	mp = map[int]int{}
	fmt.Println(mp)

}
