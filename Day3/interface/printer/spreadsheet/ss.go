package spreadsheet

import "fmt"

type SpreadSheet struct {
}

func (s SpreadSheet) Print() {
	fmt.Println("Printing a spread sheet")
}

func (s SpreadSheet) PageSetup() {
	fmt.Println("Page Setup spread sheet")
}

func (s SpreadSheet) Preview() {
	fmt.Println("Previwing Spread Sheet")
}
