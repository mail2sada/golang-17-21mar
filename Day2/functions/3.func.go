package main

import "fmt"

func main() {
	fmt.Println("Demo: functions returning value...")

	res := add(10, 20)

	fmt.Printf("\nResult is %d\n", res)

	res = add(30, 40)

	fmt.Println("Result is:", res)

	fmt.Println(sub(10, 5))
	fmt.Println(multiply(10, 5))
	fmt.Println(divide(10, 5))

	fmt.Println(divide(10, 0))

}

func add(x, y int) int {
	sum := x + y
	return sum
}

func sub(x int, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	if y == 0 {
		return -1
	}
	return x / y
}
