package main

import (
	"fmt"
)

/*
Add will compute and write result to channel
main go routine will read from channel and print results...
*/

var channel chan int = make(chan int)

func Add(a, b int) {
	sum := a + b
	channel <- sum
}

func main() {
	fmt.Println("Demo: channels")

	go Add(10, 20)

	go Add(20, 30)

	ans := <-channel
	fmt.Println(ans)

	ans = <-channel
	fmt.Println(ans)
}
