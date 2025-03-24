package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Staring server on port 1234")
	list, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Received error and can't proceed further")
		os.Exit(1)
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println("Failed to connect...")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Read n char", n, string(b))
		}
		n, err = conn.Write([]byte("Ping received suceessfully"))
		if err != nil {
			fmt.Println(err)
		}
	}
}
