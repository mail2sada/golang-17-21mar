package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Marks struct {
	M1     int     `xml:"math" json:"mathj"`
	M2     int     `xml:"science" json:"sciencej"`
	M3     int     `xml:"english" json:"english"`
	Total  int     `xml:"total" json:"total"`
	Avg    float64 `xml:"percentage" json:"percentage"`
	Result string  `xml:"grade" json:"grade"`
}

func (m *Marks) ComputeResult() {
	m.Total = (m.M1 + m.M2 + m.M3)
	m.Avg = float64(m.Total) / 3
	switch {
	case m.Avg < 50:
		m.Result = "Fail"
	case m.Avg < 60:
		m.Result = "Pass"
	case m.Avg < 70:
		m.Result = "Second"
	case m.Avg < 80:
		m.Result = "First Divison"
	default:
		m.Result = "Distinction.."
	}
}

type Student struct {
	RollNo int    `xml:"rollno" json:"rollno"`
	Name   string `xml:"name" json:"name"`
	Class  string `xml:"class" json:"class"`
	Marks  `xml:"result" json:"result"`
}

func main() {
	fmt.Println("XML Encoding...")
	list := make([]Student, 0)
	std := Student{RollNo: 1,
		Name:  "Anil",
		Class: "11",
		Marks: Marks{M1: 100,
			M2: 98,
			M3: 99}}
	std.ComputeResult()
	list = append(list, std)
	std1 := Student{RollNo: 1,
		Name:  "Anil",
		Class: "11",
		Marks: Marks{M1: 100,
			M2: 98,
			M3: 99}}
	std1.ComputeResult()
	list = append(list, std1)

	std2 := Student{RollNo: 1,
		Name:  "Anil",
		Class: "11",
		Marks: Marks{M1: 100,
			M2: 98,
			M3: 99}}
	std2.ComputeResult()
	list = append(list, std2)

	data, err := xml.Marshal(list)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	data, err = xml.MarshalIndent(list, "\t", "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	data, err = json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}
}
