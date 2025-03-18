package main

import "fmt"

type Marks struct {
	M1, M2, M3 int
}

type Student struct {
	RollNo      int
	Name, Class string
	Marks
}

func (s Student) SetRollNo(rollno int) Student {
	s.RollNo = rollno
	return s
}

func (s Student) SetName(name string) Student {
	s.Name = name
	return s
}

func (s Student) SetClass(class string) Student {
	s.Class = class
	return s
}

func main() {
	fmt.Println("Demo decorative pattern")

	s := Student{}

	s = s.SetRollNo(1).
		SetName("Anand").
		SetClass("10")
	fmt.Println(s)
}
