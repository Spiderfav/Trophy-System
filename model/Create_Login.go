package CreateLogin

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //Used for the SQL libraries I need
	"golang.org/x/crypto/bcrypt"
)

// Exporting hashing algorithm to web.go file
func hashAndSalt(pwd []byte) string {
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
	emailExist, err1 := db.Query("SELECT email FROM db.user_detail WHERE email=?", e)
	if err1 == nil {
		log.Println("Email already in db:", emailExist)
		ok = false
		message = "Email Already In DB"
		return ok, message
	} else {
		userExist, err2 := db.Query("SELECT username FROM db.user_detail WHERE username=?", u)
		if err2 == nil {
			log.Println("Username already in db:", userExist)
			ok = false
			message = "Username Already In DB"
			return ok, message
		} else {
			passBytes := []byte(p)
			passHash := hashAndSalt(passBytes)
			insertDB, err := db.Prepare("INSERT INTO user_detail(email, password, username) VALUES (?,?,?)")
			if err != nil {
				panic(err.Error())
			}
			insertDB.Exec(e, passHash, u)
			ok = true
			message = "User created"
			return ok, message

		}
	}

}

// Used in main.go
func Login(username, password string) (x bool) {
	db := dbConn()
	dbRow, err := db.Query("SELECT password FROM db.user_detail WHERE username=?", username)
	if err != nil {
		panic(err.Error())
	}
	for dbRow.Next() {
		var dbpass string
		if err := dbRow.Scan(&dbpass); err != nil {
			log.Fatal(err)
		}
		log.Print("Password DB :::>", dbpass)
		passwordByte := []byte(password)
		byteHash := []byte(dbpass)
		err3 := bcrypt.CompareHashAndPassword(byteHash, passwordByte)
		if err3 != nil {
			log.Println(err3)
			x = false
			return x
		}

		x = true
		return x

	}

	return x

}

func dbConn() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:password@tcp(192.168.0.133:3306)/db")
	if err != nil {
		panic(err.Error())
	}
	return db
}
