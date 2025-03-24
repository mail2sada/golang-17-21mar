package main

import (
	"fmt"
	"sync"
)

/*
Go routine       Goroutine					Groutine

	Reading			Processing -Squaring    Display
*/
var wg = sync.WaitGroup{}

func Generate() chan int {
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 10000; i <= 10100; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func Process(ch chan int) chan int {
	wch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			sq := val * val
			wch <- sq
		}
		close(wch)
	}()

	return wch
}

func Render(ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch {
			fmt.Println(val)
		}
	}()
}

func main() {
	fmt.Println("Pipelines...")

	// gench := Generate()
	// procch := Process(gench)
	// Render(procch)

	//Render(Process(Generate()))

	Render(Process(Process(Generate())))
	wg.Wait()
}
