package main

import (
	"fmt"
	"scope/dummy"
	"time"
)

func main() {
	fmt.Println("Exploring scope...")

	dummy.Add(10, 20)
	fmt.Println(dummy.Greet)

	var a string = "Hello"

	{
		var a string = "Wonderful.."
		fmt.Println(a) // Wondeful
	}

	fmt.Println(a)

	switch weekday := time.Now().Weekday(); weekday {
	case time.Monday:
		fmt.Println("M")
	case time.Tuesday:
		fmt.Println("T")
	case time.Wednesday:
		fmt.Println("W")
	}

	if weekday := time.Now().Weekday(); weekday == time.Wednesday {
		fmt.Println(weekday)
	}

}
