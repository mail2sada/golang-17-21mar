// package main

// import "fmt"

// type Marks struct {
// 	M1, M2, M3 int
// }

// func (m Marks) GetTotal() int {
// 	return m.M1 + m.M2 + m.M3
// }

// func (m Marks) GetAverage() float64 {
// 	return float64(m.GetTotal()) / 3
// }

// func (m Marks) PrintMarks() {
// 	fmt.Println(m.GetTotal())
// 	fmt.Println(m.GetAverage())
// }

// type Student struct {
// 	RollNo int
// 	Name   string
// 	Class  string
// 	Grade  Marks
// }

// func main() {
// 	fmt.Println("Demonstraiting methods...")

// 	std := Student{RollNo: 1,
// 		Name:  "abc",
// 		Class: "10",
// 		Grade: Marks{M1: 100,
// 			M2: 98,
// 			M3: 99}}

// 	std.Grade.PrintMarks()
// }
