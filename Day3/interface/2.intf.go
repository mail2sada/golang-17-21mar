package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Test struct {
}

func (t Test) Read() string {
	return "Testing this code..."
}

func (t Test) Write(str string) {
	fmt.Println(str)
}

func main() {

	t := Test{}

	fmt.Println(t.Read())
	t.Write("Hello..")

	var r Reader

	r = t

	fmt.Println(r.Read())

	var w Writer = t

	w.Write("Hello world..")

	var rw ReadWriter = t

	rw.Read()
	rw.Write()

}
