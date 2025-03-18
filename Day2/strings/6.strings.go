package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Lets explore strings.Contains functios")

	str := "Mavenir Systems..."

	res := strings.Contains(str, "s")

	if res {
		fmt.Println("Contains s")
	}

	res = strings.ContainsAny(str, "bc")

	if res {
		fmt.Println("Contains on of abc")
	}
}
