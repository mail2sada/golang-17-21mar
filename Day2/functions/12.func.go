package main

import "fmt"

func DemoDefer() {
	fmt.Println("Inside defer")
}

func main() {
	fmt.Println("Demo defer....")
	defer DemoDefer()
}
