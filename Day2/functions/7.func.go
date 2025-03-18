package main

import "fmt"

func main() {
	fmt.Println("Demo: Annonymous functions...")

	a := func(a string) string {
		return "Hello " + a
	}("Anand")

	fmt.Println(a)

	fmt.Println(a("Abhay"), a("Anilkumar"), a("Anand"))

}
