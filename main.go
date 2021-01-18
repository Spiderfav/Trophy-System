package main

import (
	model "Trophy-System/model"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type pageVariables struct {
	Date string
	Time string
}

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HTTPOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

/*
	This is the start of the web server with the REST API
*/
func main() {
	// http.HandleFunc() takes two inputs, the first being a pattern which is a string
	// and the second is a handler (a function that needs a ResponseWriter and a pointer to a Request).

	// Handle static files such as CSS and JS for the webpages
	fs := http.FileServer(http.Dir("view/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginOrcreate)
	http.HandleFunc("/createuser", createUser)
	http.HandleFunc("/userlogin", login)

	//http.HandleFunc("/account", homePage)  //Gonna be used later to implement the create account and login code already made
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()              // find the time right now
	homePageVars := pageVariables{ //store the date,time and username in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("view/mainpage.html") //parse the html file homepage.html
	if err != nil {                                     // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, homePageVars) //execute the template and pass it the homePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "email", Value: "", Expires: expiration}
	http.SetCookie(w, &cookie)
	/*
	   cookie, _ = r.Cookie("email")
	   	//Check if username exists
	   	if(geusernamefromdb == cookie.Value){
	   		//show Username webpage
	   	}
	*/
	w.WriteHeader(http.StatusOK)
}

func login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	match, usernameDB := model.Login(email, password)
	log.Print(match)
	log.Print(usernameDB)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "email", Value: email, Expires: expiration}
	http.SetCookie(w, &cookie)

	//cookieValue, _ := r.Cookie("email")
	//fmt.Fprint(w, cookieValue)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mapD := map[string]bool{"status": match}
	mapB, _ := json.Marshal(mapD)
	//jsonData := []byte(`{"status":match}`)
	w.Write(mapB)

}

func createUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")

	ok, message := model.CreateAccount(username, password, email)

	log.Print("Answer:", message, "=", ok) //log it

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	s := strconv.FormatBool(ok)

	mapD := map[string]string{"status": s,
		"message": message}
	mapB, _ := json.Marshal(mapD)
	w.Write(mapB)

}

func loginOrcreate(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("view/login_create.html") //parse the html file homepage.html
	if err != nil {                                         // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, 0) //execute the template and pass it the homePageVars struct to fill in the gaps
	if err != nil {       // if there is an error
		log.Print("template executing error: ", err) //log it
	}

	log.Print("I am here! Login or create")
}
