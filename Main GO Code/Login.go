package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("Let's Login to your account shall we?")
	fmt.Println("What is the username you used to login to the account? ")
	var user string
	_, err := fmt.Scan(&user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
	username, err1 := ioutil.ReadFile("Rui.txt")
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Printf("File contents: %s\n", username)
}
