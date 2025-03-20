package main

import "fmt"

type empty interface {
}

func main() {
	fmt.Println("Empty interfaces...")

	// a := 0

	// var e empty

	// e = a

	// fmt.Println(e)

	// e = "hello"

	// fmt.Println(e)

	// e = struct {
	// 	a, b int
	// }{a: 10, b: 20}

	// fmt.Println(e)

	PrintContents(10)

	PrintContents("Hello")

	PrintContents(struct {
		a, b int
	}{a: 10, b: 20})

}

func PrintContents(contents interface{}) {
	fmt.Println(contents)
}
