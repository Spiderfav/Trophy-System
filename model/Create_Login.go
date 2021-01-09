package CreateLogin

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Type CREATE to Create a User:")
	fmt.Println("Or LOGIN to Login:")
	var input string
	fmt.Scanln(&input)
	inputupper := strings.ToUpper(input)
	if inputupper == "CREATE" {
		createAccount()

	} else if inputupper == "LOGIN" {
		Login("Spiderfav")

	} else {
		fmt.Println("Not an answer")
		main()
	}
}

// Exporting hashing algorithm to web.go file
func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//func infopass(user string, pass string, email string, first string, last string) {
//}

func createAccount() {
	db := dbConn()
	fmt.Println("Make a new account.")
	fmt.Println("Username:")
	var user string
	fmt.Scanln(&user)
	fmt.Println("Password:")
	var pass string
	fmt.Scanln(&pass)
	passBytes := []byte(pass)
	passHash := HashAndSalt(passBytes)
	fmt.Println("Email:")
	var email string
	fmt.Scanln(&email)
	fmt.Println("Forename:")
	var fName string
	fmt.Scanln(&fName)
	fmt.Println("Lastname:")
	var lName string
	fmt.Scanln(&lName)
	// table value =? is the same as what im doing below
	insertDB, err := db.Prepare("INSERT INTO user_detail(email, first_name, last_name, password, username) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(email, fName, lName, passHash, user)
}

// Used in main.go
func Login(username string) (dbpass string) {
	db := dbConn()
	dbRow, err := db.Query("SELECT password FROM db.user_detail WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	for dbRow.Next() {
		//var dbpass string
		if err := dbRow.Scan(&dbpass); err != nil {
			log.Fatal(err)
		}
		log.Print("Password DB :::>", dbpass)

	}
	return dbpass

}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:password@tcp(192.168.0.133:3306)/db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
