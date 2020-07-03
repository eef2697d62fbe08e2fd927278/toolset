package user

import (
	"log"
	"net/http"
)

// HandleCreateUser : creates a user
func HandleCreateUser(res http.ResponseWriter, req *http.Request) {
	var u = NewUser("test@email.com", "username", "password")

	u.Insert()
	log.Println("Sucessfully inserted user into db")
}
