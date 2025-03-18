package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Demo: strings")

	str1 := "Mavenir "

	str2 := "Golang training"

	fmt.Println(strings.ToLower(str1 + str2))
}
