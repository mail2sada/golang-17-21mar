package main

import "fmt"

func main() {
	fmt.Println("Demo strings")

	str := "hello\nHi\"Welcome\""

	fmt.Println(str)

	fmt.Println(len(str))

	rawString := `Hello we are in "Golang training"
	We are currently working on raw strings
	92840912804!@#$$^#$%`

	fmt.Println(rawString)

}
