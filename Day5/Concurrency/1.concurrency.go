package main

import (
	"fmt"
	"time"
)

func GenerateNumber() chan int {
	ch := make(chan int)
	go func() {
		//defer close(ch)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
	}()
	return ch
}

func GenerateFloat() chan float64 {
	ch := make(chan float64)
	go func() {
		//defer close(ch)
		for i := 0.0; i < 10.0; i += 0.5 {
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
	}()

	return ch
}

func main() {
	fmt.Println("Demo:Select...")

	chi := GenerateNumber()

	chf := GenerateFloat()

	// go func() {
	// 	for i := range chi {
	// 		fmt.Println(i)
	// 	}
	// }()

	// for f := range chf {
	// 	fmt.Println(f)
	// }
	tck := time.NewTicker(5 * time.Second)
outer:
	for {
		select {
		case i := <-chi:
		
			fmt.Println(i)
			tck.Reset(5 * time.Second)
		case f := <-chf:
			tck.Reset(5 * time.Second)
			fmt.Println(f)
		case <-tck.C:
			tck.Stop()
			fmt.Println("Timedout")
			break outer
		}
	}
	close(chi)
	close(chf)
}
