package main

import "fmt"

func main() {
	fmt.Println("Annonymous structs")

	v := struct {
		a, b, c int
	}{a: 1, b: 2, c: 3}

	fmt.Println(v)

	c := v

	c.a = 10
	c.b = 20
	c.c = 30

	fmt.Println(v, c)
}
