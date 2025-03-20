package main

import "fmt"

type Student struct {
	RollNo          int
	Class, Division string
}

func (a Student) String() string {

	return fmt.Sprintln("RollNo:", a.RollNo, "\nClass:", a.Class, "\nDivision:", a.Division)
}

func main() {
	fmt.Println("Demo stringer")

	s := Student{RollNo: 1, Class: "10", Division: "B"}

	fmt.Println(s)
}
