package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = `{
	"empid":1
}`
var jsonStr1 = `{
	"name": "Anil"
}`

type Employee struct {
	EmpID int    `json:"empid"`
	Name  string `json:"name"`
	Dept  string `json:"dept"`
}

func main() {
	fmt.Println("Demo: Json encoding")
	emp := Employee{}
	err := json.Unmarshal([]byte(jsonStr), &emp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(emp)
	}

	err = json.Unmarshal([]byte(jsonStr1), &emp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(emp)
	}

	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Received error while Marshaling...", err)
	} else {
		fmt.Println(string(data))
	}

	data, err = json.MarshalIndent(emp, "\t", "")
	if err != nil {
		fmt.Println("Received error while Marshaling...", err)
	} else {
		fmt.Println(string(data))
	}
}
