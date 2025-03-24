package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo: Timer")
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timer Expired")
		wg.Done()
	})

	tch := time.After(5 * time.Second)
	tm := <-tch
	fmt.Println(tm)
	wg.Wait()

}
