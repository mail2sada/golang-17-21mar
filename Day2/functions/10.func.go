package main

import "fmt"

func main() {
	fmt.Println("Demo: defer")

	defer fmt.Println("This is defered print...")
	defer fmt.Println("This is second defered print...")

	fmt.Println("Exiting main...")
}
