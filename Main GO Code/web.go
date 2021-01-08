package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type pageVariables struct {
	Date string
	Time string
}

/*
	This is the start of the web server with the REST API
*/
func main() {
	// http.HandleFunc() takes two inputs, the first being a pattern which is a string
	// and the second is a handler (a function that needs a ResponseWriter and a pointer to a Request).
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginUser)
	//http.HandleFunc("/account", homePage)  //Gonna be used later to implement the create account and login code already made
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()              // find the time right now
	homePageVars := pageVariables{ //store the date,time and username in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, homePageVars) //execute the template and pass it the homePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("pass")

	passBytes := []byte(password)
	passHash := hashAndSalt(passBytes)

	log.Print("NAME ", username) //log it
	log.Print("PASS ", passHash) //log it

	dbRow, err := db.Query("SELECT password FROM db.user_detail WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}

	for dbRow.Next() {
		var dbpass string
		if err := dbRow.Scan(&dbpass); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Password :::>", dbpass)
	}

}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:password@tcp(192.168.0.133:3306)/db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
