package main

import "fmt"

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide by zero %d", b)
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling...")

	ans, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Received error can't proceed")
	} else {
		fmt.Println("Result is ", ans)
	}
}
