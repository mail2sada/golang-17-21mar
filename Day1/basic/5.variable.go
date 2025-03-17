package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Variable multiple declaration...")

	var a, b, c int = 10, 20, 30

	fmt.Println(a, b, c)

	var p, q, r, d = 10, 10.5, "Hello", false

	fmt.Println(p, q, r, d)

	fmt.Printf("%T, %T, %T, %T\n", p, q, r, d)

	fmt.Printf("Value of a %d, value of b %d", a, b)

	fmt.Println("Value of a", a, "Value of b", b)

	fmt.Println(reflect.TypeOf(p))
}
