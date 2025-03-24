package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg = sync.WaitGroup{}

var mutex = sync.Mutex{}

func funcA() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func funcB() {
	defer wg.Done()
	for i := 0; i < 20000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)
	fmt.Println("Demo: sync package")
	go funcA()
	go funcB()
	wg.Wait()
	fmt.Println(counter)
}
