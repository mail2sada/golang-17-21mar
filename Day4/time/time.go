package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Time")

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Add(45 * time.Minute))
	time.Sleep(2 * time.Second)
	fmt.Println(time.Since(t))

	t1 := time.Date(2025, 1, 31, 24, 10, 1, 1, &time.Location{})

	fmt.Println(time.Since(t1))

	layout := "02/01/2006 15:04:05"
	value := "20/03/2025 14:38:00"
	tx, _ := time.Parse(layout, value)
	fmt.Println(tx)

	tstr := time.Now().Format(layout)
	fmt.Println(tstr)

}
