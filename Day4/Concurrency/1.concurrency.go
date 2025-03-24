package main

import (
	"fmt"
	"time"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	go fmt.Println("We are calling a go routine...")
	fmt.Println("Working with concurrency...")
	go fmt.Println("This second go routine")

	go Add(10, 20)
	time.Sleep(1 * time.Microsecond)
}
