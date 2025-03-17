package main

import "fmt"

func main() {

	var a, b = 10, 20

	var c = a + b

	fmt.Println(c)

	c += a + b

	fmt.Println(c)

	c -= a + b
	fmt.Println(c)

	c *= 2
	fmt.Println(c)

}
