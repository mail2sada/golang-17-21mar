package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 3)

func Write() {
	i := 0
	for {
		time.Sleep(10 * time.Millisecond)
		i++
		ch <- i
	}
}

func Read() {
	time.Sleep(1 * time.Second)
	for {
		time.Sleep(10 * time.Millisecond)
		val := <-ch
		fmt.Println(val)
	}
}

func main() {
	fmt.Println("Demo buffered channels")
	go Write()
	go Read()
	for {
		fmt.Println("cap:", cap(ch))
		fmt.Println("len:", len(ch))
		time.Sleep(30 * time.Millisecond)
	}
}
