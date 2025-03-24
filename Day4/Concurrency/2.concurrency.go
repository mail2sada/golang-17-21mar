package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Demo: Wait groups")
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("First go routine...")

	}()
	//wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Second go routine...")

	}()
	//wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Third go routine...")

	}()

	wg.Wait()
}
