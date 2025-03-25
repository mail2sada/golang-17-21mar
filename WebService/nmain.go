package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
}

func main() {
	list := []Student{} // empty slice

	s1 := Student{RollNo: 1, Name: "A"}
	s2 := Student{RollNo: 2, Name: "B"}

	list = append(list, s1)

	list = append(list, s2)

	fmt.Println(list)
}
