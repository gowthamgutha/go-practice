package main

import (
	"fmt"
	"net/http"
)

func hello(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello world!")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/headers", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(rw, "User agent", req.UserAgent())
	})

	http.ListenAndServe(":8090", nil)

}
