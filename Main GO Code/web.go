package main

import (
	"fmt"
	"net/http"
)

/*
	This is the start of the web server with the REST API
*/
func main() {
	http.HandleFunc("/login", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}
