package user

import (
	"log"
	"net/http"
)

// CreateUser : creates a user
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var u = NewUser("test@email.com", "username", "password")

	u.Insert()
	log.Println("Sucessfully inserted user into db")
}

