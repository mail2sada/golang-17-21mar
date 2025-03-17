package main

import "fmt"

type Marks struct {
	RollNo     int
	M1, M2, M3 int
	Total      int
	Average    float64
}

type Student struct {
	RollNo                int
	Name, Class, Division string
	Grades                Marks
}

func main() {
	fmt.Println("Nested structures...")

	s1 := Student{RollNo: 1,
		Name: "abcdef", Class: "10",
		Division: "A",
		Grades: Marks{RollNo: 1,
			M1: 100,
			M2: 85,
			M3: 95}}
	fmt.Println(s1)
}
