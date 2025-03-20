package main

import (
	"exploration/greet"
	_ "exploration/init"
	mt "exploration/math"
	"exploration/math/advanced"
	"exploration/math/geometry"
	"fmt"
	"math"

	"github.com/fatih/color"
)

func init() {
	fmt.Println("This is my init")
}

func init() {
	fmt.Println("init 2")
}

func init() {
	fmt.Println("init 3")
}

func main() {
	fmt.Println("Exploring packages and modules..")
	res := mt.Add(10, 20)
	fmt.Println(res)
	res = mt.Sub(20, 10)
	fmt.Println(res)

	circle := geometry.Circle{Radius: 2.4}

	area := circle.GetArea()
	fmt.Println(area)
	circ := circle.GetCircumfrance()
	fmt.Println(circ)

	rect := geometry.Rectangle{Lenght: 10, Breadth: 5}

	area = rect.GetArea()
	fmt.Println(area)
	perim := rect.GetPerimeter()
	fmt.Println(perim)

	fmt.Println(advanced.Square(10))

	fmt.Println(advanced.Sin(90))

	fmt.Println(greet.GreetMessage("Bharat"))
	color.Cyan("Using color package github.com/fatih/color")

	color.RGB(100, 200, 255).Println("Hello.")

	fmt.Println(math.Abs(-12.344))

}

func init() {
	fmt.Println("init 4")
}

func init() {
	fmt.Println("init 5")
}
