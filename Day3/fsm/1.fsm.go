package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Working with directories")

	//Create Directory

	err := os.Mkdir("Test", 0677)

}
