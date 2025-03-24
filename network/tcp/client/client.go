package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Demo: tcp client...")
	conn, err := net.Dial("tcp", ":1234")

	if err != nil {
		fmt.Println("Received error", err)
		os.Exit(1)
	}
	for {
		n, err := conn.Write([]byte("Ping from client...."))
		if err != nil {
			fmt.Println("received error", err)
		} else {
			fmt.Println("Wrote n bytes...", n)
		}
		data := make([]byte, 1024)
		n, err = conn.Read(data)
		if err != nil {
			fmt.Println("received error", err)
		} else {
			fmt.Println(string(data))
		}
	}
}
