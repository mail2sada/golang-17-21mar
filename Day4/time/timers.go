package main

import (
	"fmt"
	"time"
)

var cch chan chan time.Time

func trackTimers(ch <-chan time.Time) {
	cch <- ch
}

func handleTimers(cch chan chan time.Time) {
	go func() {
		for val := range cch {
			tm := <-val
			fmt.Println(tm)
		}
	}()
}

func main() {

	cch = make(chan chan time.Time)

	trackTimers(time.After(1 * time.Second))

	trackTimers(time.After(2 * time.Second))

	trackTimers(time.After(3 * time.Second))
}
