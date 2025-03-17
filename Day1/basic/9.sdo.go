package main

import "fmt"

func main() {
	fmt.Println("Demo: short declaration operator")

	a := 10

	fmt.Println(a)

	b, c := 10, 20

	fmt.Println(b, c)

	e, f := "Hello", false

	fmt.Println(e, f)
}
