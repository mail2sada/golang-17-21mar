package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch case...")

	var a = "Monday"

	switch a {
	case "Monday":
		fallthrough
	case "Tuesday":
		fallthrough
	case "Wensday":
		fmt.Println("Weekday")
	case "Saturday":
		fallthrough
	case "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Unknown")
	}
}
