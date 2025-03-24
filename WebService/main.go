package main

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	serviceurl = ":3456"
)

func genericHttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	if strings.Contains(r.RequestURI, "user") {
		// this will handled by user
		fmt.Println("user")
	} else if strings.Contains(r.RequestURI, "product") {
		// this will handled by user
		fmt.Println("product")
	} else if strings.Contains(r.RequestURI, "cart") {
		// this will handled by user
		fmt.Println("cart")
	} else {
		fmt.Println("root route..")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("We have received request", r.URL)
}

var (
	serviceRoutes = []string{"/", "/user", "/product", "/cart"}
)

func createServiceRoute() {
	for _, v := range serviceRoutes {
		http.HandleFunc(v, genericHttpHandler)
	}
}

func main() {
	fmt.Println("Simple web service..")
	createServiceRoute()
	http.ListenAndServe(serviceurl, nil)
}
