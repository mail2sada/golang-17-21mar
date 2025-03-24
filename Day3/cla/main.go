package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Demo commandline arguments...")
	// fmt.Println("Arguments passed to the program ", os.Args)

	// args := os.Args[1:]

	// for _, param := range args {
	// 	fmt.Println(param)
	// }

	/*
		string case converter

		./cla -case upper -arg sdhsjdhks
	*/

	c := flag.String("case", "upper", "usage: coverts to the specified case possible values are lower & upper")

	a := flag.String("arg", "", "converts this value to a specified case in -case parameter")

	flag.StringVar()

	flag.Parse()

	switch *c {
	case "upper":
		fmt.Println(strings.ToUpper(*a))
	case "lower":
		fmt.Println(strings.ToLower(*a))
	case "title":
		fmt.Println(strings.ToTitle(*a))
	default:
		flag.Usage()
	}
}
