package main

import (
	"fmt"
)

type CustomError struct {
	errorCode int
	errorMsg  string
}

func (c CustomError) Error() string {
	return fmt.Sprintf("\nError Code %d and Error Message %s\n", c.errorCode, c.errorMsg)
}

func getFromSlice(slc []int, index int) (ans int, err error) {
	if index >= len(slc) {
		err = CustomError{errorCode: 10, errorMsg: "Index out of bounds"}
		return
	}

	if index < 0 {
		err = CustomError{errorCode: 20, errorMsg: "Negetive index"}
		return
	}
	ans = slc[index]
	return
}

func main() {
	fmt.Println("Demo: Custom error")

	ans, err := getFromSlice([]int{12, 3, 4, 5, 6, 7, 7, 8, 9, 0}, -10)

	if err != nil {
		cerr := err.(CustomError)
		switch cerr.errorCode {
		case 10:
			fmt.Println("Out of bounds")
		case 20:
			fmt.Println("-ve index...")
		}

	} else {
		fmt.Println(ans)
	}

}
