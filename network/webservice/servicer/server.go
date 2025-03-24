package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	port = 3456
)

//var service_routes = []string{"/", "/user", "/product", "/cart"}

func main() {
	fmt.Println("Web service...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received a request on root service route...")
		fmt.Println("Request ", r)
		//fmt.Println("Method:", r.Method)
		fmt.Println("Headers:", r.Header)

		for key, val := range r.Header {
			fmt.Println(key, ":", val)
		}

		rb := r.Body
		str := ""
		for {
			body := make([]byte, 1024)
			_, err := rb.Read(body)
			if err != nil {
				if err == io.EOF {
					fmt.Println("End of reading request body")
					break
				}
				fmt.Println("Received error", err)
				return
			}
			str += string(body)

		}
		fmt.Println("Read", "bytes. Content:", str)

		for key, val := range r.URL.Query() {
			fmt.Println(key, val)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Successfully handled the request..."))
	})

	err := http.ListenAndServe(":3456", nil)
	if err != nil {
		fmt.Println("Received error can not proceed further...")
		os.Exit(1)
	}
}
