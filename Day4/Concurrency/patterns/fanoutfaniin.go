package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Task1->Taks2->Task3->Task4
	   1     4      1      1

	           Task2
               Task2
	   Taslk1  Task2  Task3  Task4
	   	  1    Task2
*/

var wg = sync.WaitGroup{}

func Generate() chan int {
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 100; i < 200; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Process(ch chan int) chan int {
	wch := make(chan int)

	fanoutCount := 4

	lwg := sync.WaitGroup{}
	for i := 0; i < fanoutCount; i++ {
		lwg.Add(1)
		go func() {
			defer lwg.Done()

			for event := range ch {
				time.Sleep(4 * time.Second)
				wch <- event * event
			}
			//close(wch)
		}()
	}

	go func() {
		lwg.Wait()
		close(wch)
	}()

	return wch
}

func Display(ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for event := range ch {
			time.Sleep(1 * time.Second)
			fmt.Println(event)
		}
	}()
}

func main() {
	fmt.Println("Demo: Fanout FanIn")

	Display(Process(Generate()))
	wg.Wait()
}
