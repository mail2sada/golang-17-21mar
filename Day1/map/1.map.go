package main

import "fmt"

func main() {
	fmt.Println("Demo maps")

	mp := map[int]int{1: 100, 2: 200, 3: 1100}

	fmt.Println(mp)

	mp[4] = 10000
	fmt.Println(mp)

	v := mp[3]
	fmt.Println(v)

	v, ok := mp[100]

	if !ok {
		fmt.Println("Value not available...")
	} else {
		fmt.Println(v)

	}

}
