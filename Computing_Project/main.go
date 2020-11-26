package main

import (
	"crypto/md5"
	"fmt"
  "encoding/hex"
  "strings"
	"database/sql"
	//"log"
	//"net/http"
	//"text/template"

	_ "github.com/go-sql-driver/mysql"
)

/*
type User struct{
	username string
	password string

}
*/


func main() {
  fmt.Println("Type CREATE to Create a User:")
  fmt.Println("Or LOGIN to Login:")
  var input string
  fmt.Scanln(&input)
  inputupper := strings.ToUpper(input)
  if inputupper == "CREATE"{
		create_account()

  }else if inputupper == "LOGIN"{
		login()


  }else{
    fmt.Println("Not an answer")
		main()
  }
}

func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}

func infopass(user string, pass string, email string, firs string, last string){
}

func create_account(){
	db := dbConn()
	fmt.Println("Make a new account.")
	fmt.Println("Username:")
	var user string
	fmt.Scanln(&user)
	fmt.Println("Password:")
	var pass string
	fmt.Scanln(&pass)
	pass_hash := GetMD5Hash(pass)
	fmt.Println("Email:")
	var email string
	fmt.Scanln(&email)
	fmt.Println("Forename:")
	var f_name string
	fmt.Scanln(&f_name)
	fmt.Println("Lastname:")
	var l_name string
	fmt.Scanln(&l_name)
	// table value =? is the same as what im doing below
	insertDB, err := db.Prepare("INSERT INTO user_detail(email, first_name, last_name, password, username) VALUES (?,?,?,?,?)")
	if err != nil {
			panic(err.Error())
	}
	insertDB.Exec(email,f_name,l_name,pass_hash,user)
}

func login(){
	db := dbConn()
	fmt.Println("Username:")
	var user string
	fmt.Scanln(&user)
	fmt.Println("Password:")
	var pass string
	fmt.Scanln(&pass)
	pass_hash := GetMD5Hash(pass)
	dbRow, err := db.Query("SELECT password FROM db.user_detail WHERE username=?", user)
	if err != nil {
			panic(err.Error())
			fmt.Println("User not found.")
	}

	if checkDB != pass_hash{
		fmt.Println("Wrong password.")
	}

}

func dbConn() (db *sql.DB) {
    db, err := sql.Open("mysql","user:password@tcp(192.168.0.32:3306)/db")
    if err != nil {
        panic(err.Error())
    }
    return db
}
