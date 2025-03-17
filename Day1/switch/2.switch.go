package main

import "fmt"

func main() {
	fmt.Println("Demo: Switch..")

	a := 10

	switch a {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Unhandled.. ")

	}
}
