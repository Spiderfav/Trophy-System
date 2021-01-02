package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

func main1() {
	db := dbConn()
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

	// table value =? is the same as what im doing below
	insertDB, err := db.Prepare("INSERT INTO user_detail(email, first_name, last_name, password, username) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(email, f_name, l_name, pass_hash, user)
}

// Function made by me previously, connects to the DB I am using on the Raspi
func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:password@tcp(192.168.0.133:3306)/db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
