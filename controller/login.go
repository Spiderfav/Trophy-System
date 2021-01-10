package login

import (
	model "A-Level-Trophy-System/model"
	"log"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {

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
