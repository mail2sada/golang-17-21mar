package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int)
var wg = sync.WaitGroup{}

func Test() {
	defer wg.Done()

	fmt.Println("Testing blocking reads and writes of go channels")
	// write
	fmt.Println("Writing...")
	time.Sleep(1 * time.Second)
	ch <- 10
	fmt.Println("Wrote 10 to ch")

}

func Best() {
	defer wg.Done()
	fmt.Println("We are in best")
	fmt.Println("Reading")
	s := <-ch
	fmt.Println("Read value from ch", s)
}

func main() {
	wg.Add(2)
	fmt.Println("Demo: Channels...")

	go Test()
	go Best()

	wg.Wait()
}
