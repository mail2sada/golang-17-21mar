package main

import "fmt"

type Std struct {
	RollNo      int
	Name, Class string
	Grade       struct {
		M1, M2, M3 int
		Total      int
		Avg        float64
	}
}

func main() {
	fmt.Println("Demo: Annonymous structures")

	std1 := Std{
		RollNo: 1,
		Name:   "Anand",
		Class:  "10",
		Grade: struct {
			M1    int
			M2    int
			M3    int
			Total int
			Avg   float64
		}{M1: 10, M2: 20, M3: 30},
	}
	fmt.Println(std1)

	std1.Grade.Total = std1.Grade.M1 + std1.Grade.M2 + std1.Grade.M3
	std1.Grade.Avg = float64(std1.Grade.Total) / 3
	fmt.Println(std1)

}
