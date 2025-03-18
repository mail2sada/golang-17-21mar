package main

import "fmt"

func main() {
	fmt.Println("String comparison")

	str := "test"

	str1 := "test1"

	if str == str1 {
		fmt.Println("Strings are equal...")
	}

	if str != str1 {
		fmt.Println("Strings are not equal...")
	}

}
