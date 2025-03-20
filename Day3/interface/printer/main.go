package main

import (
	"fmt"
	"priniter/document"
	"priniter/print"
	"priniter/spreadsheet"
)

func main() {
	fmt.Println("Demo: Printer interfaces...")

	d := document.Document{}

	s := spreadsheet.SpreadSheet{}

	print.Print(d)

	print.PageSetup(s)
}
