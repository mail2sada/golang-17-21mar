package main

import (
	"fmt"
	"sync"
	"time"
)

func handleTicker(tc *time.Ticker, wg *sync.WaitGroup) {
	defer wg.Done()
	for event := range tc.C {
		fmt.Println("Ticker fired:", event)
	}
}

const (
	RESET = 0
	STOP  = 1
)

func tickerHandler(tck *time.Ticker, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for {
		time.Sleep(2 * time.Second)
		if count < 5 {
			tck.Reset(500 * time.Millisecond)
		}
		if count == 5 {
			tck.Stop()
			break
		}
		count++
	}
}

func main() {
	fmt.Println("Demo: Ticker...")
	wg := sync.WaitGroup{}
	wg.Add(2)
	tcr := time.NewTicker(1 * time.Second)
	go handleTicker(tcr, &wg)

	go tickerHandler(tcr, &wg)
	wg.Wait()
}
