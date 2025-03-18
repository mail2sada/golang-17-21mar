package main

import "fmt"

func AddNumbers(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Call by value and reference")

	x, y := 100, 200

	c := AddNumbers(x, y)

	fmt.Println(c)

	c = CallWithRef(&x, &y)

	fmt.Println(c)

	fmt.Println(x, y)

}

func CallWithRef(a, b *int) int {
	sum := *a + *b
	*a = 100000
	*b = 200000
	return sum
}
