package main

import "fmt"

func main() {
	fmt.Println("Demo: Function is with parameters")

	greetWithName("Anil kumar")

	greetWithName("Anand singh")

	greetWithName("Arun kumar")
}

func greetWithName(name string) {
	fmt.Println("Welcome ", name)
}
