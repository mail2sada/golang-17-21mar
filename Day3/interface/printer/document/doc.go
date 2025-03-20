package document

import "fmt"

type Document struct {
}

func (d Document) Print() {
	fmt.Println("Printing a document")
}

func (d Document) Preview() {
	fmt.Println("Previewing a document..")
}

func (d Document) PageSetup() {
	fmt.Println("Document PageSetup")
}
