package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Let's create an account shall we?")
	fmt.Println("What is the username you want to use? ")
	var user string
	_, err := fmt.Scan(&user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
	userbyte := []byte(user)
	err = ioutil.WriteFile("Rui.txt", userbyte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
