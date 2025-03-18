package main

import "fmt"

type DemoStruct struct {
	int
	string
	float32
	float64
	bool
}

func main() {
	fmt.Println("Demo.. struct as parameter")

	var d DemoStruct

	ByValue(d)

	fmt.Println(d)

	ByReference(&d)

	fmt.Println(d)

}

func ByValue(a DemoStruct) {
	a.int = 100
	a.bool = true
}

func ByReference(a *DemoStruct) {
	a.float32 = 3.14
	a.float64 = 10009.0090
	a.string = "Somthing"
}
