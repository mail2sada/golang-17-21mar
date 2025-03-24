package main

import "fmt"

type Marks struct {
	M1, M2, M3 int
}

func (m Marks) GetTotal() int {
	return m.M1 + m.M2 + m.M3
}

func (m Marks) GetAverage() float64 {
	return float64(m.GetTotal()) / 3
}

func (m Marks) PrintMarks() {
	fmt.Println(m.GetTotal())
	fmt.Println(m.GetAverage())
}

type Student struct {
	RollNo int
	Name   string
	Class  string
	Marks
}

func (s *Student) UpdateRollNo(RollNo int) {
	s.RollNo = RollNo
}

func (s *Student) UpdateName(name string) {
	s.Name = name
}

func main() {
	fmt.Println("Demonstraiting methods...")

	std := Student{}

	std.UpdateRollNo(1)
	std.UpdateName("Hello..")

	fmt.Println(std)

	std.PrintMarks()

}
