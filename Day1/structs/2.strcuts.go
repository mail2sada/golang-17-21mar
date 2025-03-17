package main

import "fmt"

type Employee struct {
	Eid         int
	Name        string
	Department  string
	Designation string
	Manger      int
}

func main() {
	fmt.Println("Demo: Structus")

	var e1 = Employee{Eid: 10,
		Name:        "Prashanth",
		Department:  "NGN",
		Designation: "Manager",
		Manger:      100}

	fmt.Println(e1)

	fmt.Println("Eid:", e1.Eid, "Name:", e1.Name, "Dept:", e1.Department)

	e1.Eid = 500
	fmt.Println("Eid:", e1.Eid, "Name:", e1.Name, "Dept:", e1.Department)

}
