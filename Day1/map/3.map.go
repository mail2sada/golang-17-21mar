package main

import "fmt"

func main() {
	fmt.Println("Map iterators")

	mp := make(map[int]string)

	mp[0] = "Hello"

	mp[1] = "hi"

	mp[2] = "Mavenir"

	mp[3] = "abc"

	mp[100] = "xyz"

	fmt.Println(mp)

	for key, val := range mp {
		fmt.Println(key, val)
	}
}
