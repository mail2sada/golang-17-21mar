package main

import "fmt"

func main() {
	fmt.Println("Demo:String operations")

	str := "hi"

	str1 := "hello"

	str2 := str + str1

	fmt.Println(str2)

	str += str1

	fmt.Println(str)
}
