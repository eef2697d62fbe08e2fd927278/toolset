package user

import (
	"log"
	"net/http"
)

// HandleCreate : creates a user
func HandleCreate(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		for _, i := range req.Header["Content-Type"] {
			if i == "application/json" {
				u := NewUser("mail@mail.com", "username", "password")
				u.Insert()

				res.WriteHeader(201)
			}
		}

		res.WriteHeader(400)
	}
	var u = NewUser("test@email.com", "username", "password")

	u.Insert()
	log.Println("Sucessfully inserted user into db")
}
