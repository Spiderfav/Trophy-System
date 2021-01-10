package CreateLogin

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

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

func CreateAccount(u, p, e string) (ok bool, message string) {
	db := dbConn()
	// table value =? is the same as what im doing below
	insertDB, err := db.Prepare("INSERT INTO user_detail(email, password, username) VALUES (?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(e, p, u)
	ok = true
	message = "User created"
	return ok, message
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
