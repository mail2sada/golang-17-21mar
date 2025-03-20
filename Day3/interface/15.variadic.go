package main

import "fmt"

func VariadicPrint(collection ...interface{}) {
	for _, ele := range collection {
		fmt.Println(ele)
	}
}

func main() {
	fmt.Println("Variadic with empty interfaces...")

	VariadicPrint(1, "Hello", 10.2)
}
