package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	model "A-Level-Trophy-System/model"

	_ "github.com/go-sql-driver/mysql"
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

	t, err := template.ParseFiles("view/homepage.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, homePageVars) //execute the template and pass it the homePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("pass")

	passBytes := []byte(password)
	passHash := model.HashAndSalt(passBytes)

	log.Print("NAME ", username) //log it

	dbpass := model.Login(username)

	log.Print("PASS ", passHash)
	log.Print("PASS DB ", dbpass)

}
