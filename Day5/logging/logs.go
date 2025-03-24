package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	f, err := os.Create("debug.log")
	if err != nil {
		fmt.Println("Received error while creating file..")
		log.Fatalln("Can't proceed, received")
	}
	defer f.Close()

	fmt.Println("Demo: Logging")
	log.Println("Inside main...")
	fmt.Println("Processing")
	log.Println("Exiting main")
	info, _ := os.Create("info.log")
	defer info.Close()
	dlogger := log.New(f, "[DEBUG]", log.Ldate|log.Lshortfile)
	dlogger.Println("Printing debug log..")
	ilogger := log.New(info, "[Info]", log.Ldate|log.Lshortfile)
	ilogger.Println("Printing Info log..")

}
