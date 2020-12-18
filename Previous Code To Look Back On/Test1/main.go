package main

import (
  "fmt"
  "github.com/gorilla/mux")

func main() {
	fmt.Printf("hello, world\n")
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  //http.Handle("/", r)
}

func HomeHandler() {
  fmt.Printf("hello")
}
