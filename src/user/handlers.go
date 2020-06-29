package user

import (
	"fmt"
	"net/http"
)

// CreateUser : creates a user
func CreateUser(res http.ResponseWriter, req *http.Request) {
	var u = NewUser("test@email.com", "username", "password")

	u.Insert()
	fmt.Println("Sucessfully inserted user into db")
}
