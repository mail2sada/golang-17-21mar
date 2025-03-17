package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch case... ")

	day := "Monday"

	switch day {
	case "Saturaday":
		fmt.Println("Weekend")
	case "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Working day..")
	}
}
