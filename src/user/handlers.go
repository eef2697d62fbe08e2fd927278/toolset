package user

import (
	"log"
	"net/http"

	"google.golang.org/appengine/log"
)

// CreateUser : creates a user
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var u = NewUser("test@email.com", "username", "password")

	u.Insert()
	log.Done("Sucessfully inserted user into db")
}
