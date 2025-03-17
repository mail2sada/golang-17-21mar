package main

import "fmt"

func main() {
	fmt.Println("Demo: structures")
	type TestStruct struct {
		a, b, c string
	}

	var v = TestStruct{a: "Hi", b: "Hello", c: "Hi"}

	fmt.Println(v)
}
