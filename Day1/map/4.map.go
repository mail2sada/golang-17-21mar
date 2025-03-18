package main

import "fmt"

type temp struct {
	int
	float32
}

func main() {
	fmt.Println("Map of structures...")

	mp := make(map[int]temp)

	mp[100] = temp{int: 10, float32: 10.2}

	mp[200] = temp{int: 100, float32: 100.2}

	fmt.Println(mp)

	for _, v := range mp {
		fmt.Println(v.int, v.float32)
	}
}
